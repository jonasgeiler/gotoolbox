package gotoolbox

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
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
	binCachePathName := t.Name
	binCachePath := filepath.Join(binCacheDir, binCachePathName)
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

	binDownloadFilePattern := ".temp*-" + binDownloadFileName
	binDownloadFile, err := os.CreateTemp(
		binCacheDir,
		binDownloadFilePattern,
	)
	if err != nil {
		return "", fmt.Errorf(
			"creating temporary file with pattern %q in %q: %w",
			binDownloadFilePattern, binCacheDir, err,
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

	var finalBinDownloadFile *os.File
	if binDownloadInfo.ExtractFile == "" {
		// We have directly downloaded a binary and can just finalize it.
		finalBinDownloadFile = binDownloadFile
	} else {
		// We have downloaded an archive and need to extract it.
		archiveBinReader, err := OpenFileInArchive(
			binDownloadFile,
			binDownloadSize,
			binDownloadInfo.ExtractFile,
		)
		if err != nil {
			return "", err
		}
		defer archiveBinReader.Close()

		binFilePattern := ".temp*-" + binCachePathName
		binFile, err := os.CreateTemp(
			binCacheDir,
			binFilePattern,
		)
		if err != nil {
			return "", fmt.Errorf(
				"creating temporary file with pattern %q in %q: %w",
				binFilePattern, binCacheDir, err,
			)
		}
		defer os.Remove(binFile.Name())
		defer binFile.Close()
		if _, err = io.Copy(binFile, archiveBinReader); err != nil {
			return "", fmt.Errorf(
				"extracting file from %q to %q: %w",
				binDownloadFileName, binCachePath, err,
			)
		}

		finalBinDownloadFile = binFile
	}

	if err := finalBinDownloadFile.Chmod(0755); err != nil {
		return "", fmt.Errorf(
			"changing file permissions for %q: %w",
			finalBinDownloadFile.Name(), err,
		)
	}
	if err := os.Rename(finalBinDownloadFile.Name(), binCachePath); err != nil {
		return "", fmt.Errorf(
			"moving temporary file %q to %q: %w",
			finalBinDownloadFile.Name(), binCachePath, err,
		)
	}

	return binCachePath, nil
}
