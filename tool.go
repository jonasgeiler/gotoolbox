package gotoolbox

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// BuildType differentiates between dynamically and statically linked builds of
// release binaries.
type BuildType int

const (
	// BuildTypeStatic means that platform environment matching is NOT needed.
	BuildTypeStatic BuildType = iota

	// BuildTypeDynamic means that platform environment matching is needed.
	BuildTypeDynamic
)

// DownloadInfo describes from where to download a binary and what to do with
// it after (verify checksum, extract, etc.).
type DownloadInfo struct {
	URL         string
	Checksum    string
	ExtractFile string
}

// Tool defines a gotoolbox tool with its download and versioning info.
type Tool struct {
	Name      string
	Version   string
	BuildType BuildType
	Binaries  map[Platform]DownloadInfo
}

func (t *Tool) DownloadAndExec() {
	binPath, err := t.downloadIfNeeded()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Download Error: %v\n", err)
		os.Exit(1)
	}

	Exec(binPath)
}

func (t *Tool) downloadIfNeeded() (string, error) {
	// If the tool builds dynamically linked binaries, we need the host platform
	// with it's detected environment.
	hostPlatform := HostPlatform(t.BuildType == BuildTypeDynamic)

	binCacheDir := filepath.Join(
		ToolCacheDir(), // TODO: Use Go workspace/module root instead? Traverse upwards until a go.work file found, if not found try again but look for go.mod. This would also allow just deleting old cached versions maybe.
		t.Name,
		t.Version,
		hostPlatform.OS,
		hostPlatform.Arch,
		hostPlatform.Env.DirName(),
	)
	binCachePath := filepath.Join(binCacheDir, t.Name)
	if _, err := os.Stat(binCachePath); err == nil {
		// Already exists, skip downloading.
		return binCachePath, nil
	}

	binDownloadInfo, ok := t.Binaries[hostPlatform]
	if !ok {
		return "", fmt.Errorf(
			"no binary found for host platform: %s",
			hostPlatform,
		)
	}

	binDownloadFileName := path.Base(binDownloadInfo.URL)
	if binDownloadFileName == "." || binDownloadInfo.URL == "/" {
		return "", fmt.Errorf(
			"failed to determine download file name from %q",
			binDownloadInfo.URL,
		)
	}

	if err := os.MkdirAll(binCacheDir, 0755); err != nil {
		return "", fmt.Errorf(
			"creating binary cache directory %q: %w",
			binCacheDir, err,
		)
	}

	binDownload, err := Download(binDownloadInfo.URL)
	if err != nil {
		return "", fmt.Errorf(
			"downloading file from %q: %w",
			binDownloadInfo.URL, err,
		)
	}
	defer binDownload.Close()

	binDownloadFile, err := os.CreateTemp(binCacheDir,
		".temp*-"+binDownloadFileName,
	)
	if err != nil {
		return "", fmt.Errorf(
			"creating temporary file with pattern \".temp*-%s\": %w",
			binDownloadFileName, err,
		)
	}
	defer os.Remove(binDownloadFile.Name())
	defer binDownloadFile.Close()

	binDownloadHash := sha256.New()
	binDownloadSize, err := io.Copy(
		io.MultiWriter(binDownloadFile, binDownloadHash),
		binDownload,
	)
	if err != nil {
		return "", fmt.Errorf(
			"streaming/copying response body from %q: %w",
			binDownloadInfo.URL, err,
		)
	}
	binDownloadHashSum := hex.EncodeToString(binDownloadHash.Sum(nil))

	if binDownloadHashSum != binDownloadInfo.Checksum {
		return "", fmt.Errorf(
			"checksum mismatch for file downloaded from %q: expected %q, got %q",
			binDownloadInfo.URL, binDownloadInfo.Checksum, binDownloadHashSum,
		)
	}

	if binDownloadInfo.ExtractFile == "" {
		// We have directly downloaded a binary and can just move the temporary
		// file to its final location.
		if err := binDownloadFile.Chmod(0755); err != nil {
			return "", fmt.Errorf(
				"changing file permissions for %q: %w",
				binDownloadFile.Name(), err,
			)
		}
		if err := os.Rename(binDownloadFile.Name(), binCachePath); err != nil {
			return "", fmt.Errorf(
				"moving temporary file %q to %q: %w",
				binDownloadFile.Name(), binCachePath, err,
			)
		}
	} else {
		// We have downloaded an archive and need to extract it.
		var archiveBinReader io.Reader
		if strings.HasSuffix(binDownloadFileName, ".tar.gz") {
			if _, err = binDownloadFile.Seek(0, io.SeekStart); err != nil {
				return "", fmt.Errorf(
					"seeking to beginning of archive temp file: %w",
					err,
				)
			}
			gzipReader, err := gzip.NewReader(binDownloadFile)
			if err != nil {
				return "", fmt.Errorf(
					"creating gzip reader: %w",
					err,
				)
			}
			defer gzipReader.Close()
			tarReader := tar.NewReader(gzipReader)
			for {
				header, err := tarReader.Next()
				if err == io.EOF {
					return "", fmt.Errorf(
						"file %q not found in archive %q",
						binDownloadInfo.ExtractFile, binDownloadFileName,
					)
				}
				if err != nil {
					return "", fmt.Errorf(
						"reading tar archive %q: %w",
						binDownloadFileName, err,
					)
				}
				if header.Typeflag == tar.TypeReg && header.Name == binDownloadInfo.ExtractFile {
					// Found the binary, so exit loop.
					break
				}
			}
			archiveBinReader = tarReader
		} else if strings.HasSuffix(binDownloadFileName, ".zip") {
			zipReader, err := zip.NewReader(binDownloadFile, binDownloadSize)
			if err != nil {
				return "", fmt.Errorf("creating zip reader: %w", err)
			}
			archiveBinFile := func() *zip.File {
				for _, file := range zipReader.File {
					if file.Name == binDownloadInfo.ExtractFile {
						return file
					}
				}
				return nil
			}()
			if archiveBinFile == nil {
				return "", fmt.Errorf(
					"binary file %q not found in archive %q",
					binDownloadInfo.ExtractFile, binDownloadFileName,
				)
			}
			archiveBinReadCloser, err := archiveBinFile.Open()
			if err != nil {
				return "", fmt.Errorf(
					"opening file %q in zip archive: %w",
					archiveBinFile.Name, err,
				)
			}
			defer archiveBinReadCloser.Close()
			archiveBinReader = archiveBinReadCloser
		} else {
			return "", fmt.Errorf(
				"unsupported archive file format: %q",
				binDownloadFileName,
			)
		}
		binFile, err := os.OpenFile(
			binCachePath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			0755,
		)
		if err != nil {
			return "", fmt.Errorf(
				"creating binary file %q: %w",
				binCachePath, err,
			)
		}
		defer binFile.Close()
		if _, err = io.Copy(binFile, archiveBinReader); err != nil {
			return "", fmt.Errorf(
				"extracting file from %q to %q: %w",
				binDownloadFileName, binCachePath, err,
			)
		}
	}

	return binCachePath, nil
}
