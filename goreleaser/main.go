package main

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/jonasgeiler/gotoolbox"
)

const tool = "goreleaser"

const (
	// renovate: datasource=github-releases depName=dprint/dprint
	version    = "v2.16.0"
	repository = "goreleaser/goreleaser"
)

func main() {
	binPath, err := downloadIfNeeded()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Download Error: %v\n", err)
		os.Exit(1)
	}

	gotoolbox.Exec(binPath)
}

func downloadIfNeeded() (string, error) {
	platform, err := goreleaserPlatform()
	if err != nil {
		return "", fmt.Errorf(
			"determining %s platform: %w",
			tool, err,
		)
	}
	archiveExt := ".tar.gz"
	binExt := ""
	if runtime.GOOS == "windows" {
		archiveExt = ".zip"
		binExt = ".exe"
	}
	binPath := gotoolbox.ResolveToolPath(
		fmt.Sprintf(
			"%s-%s-%s%s",
			tool, platform, version, binExt,
		),
	)
	if _, err = os.Stat(binPath); err == nil {
		// Already exists, skip downloading.
		return binPath, nil
	}

	archiveName := fmt.Sprintf("goreleaser_%s%s", platform, archiveExt)
	archiveStream, err := gotoolbox.FetchGitHubReleaseAsset(
		repository, version, archiveName,
	)
	if err != nil {
		return "", err
	}
	defer archiveStream.Close()
	archiveTemp, err := os.CreateTemp("", "*-"+archiveName)
	if err != nil {
		return "", fmt.Errorf(
			"creating temporary file with pattern \"*-%s\": %w",
			archiveName, err,
		)
	}
	defer os.Remove(archiveTemp.Name())
	defer archiveTemp.Close()
	archiveHash := sha256.New()
	archiveSize, err := io.Copy(
		io.MultiWriter(archiveTemp, archiveHash),
		archiveStream,
	)
	if err != nil {
		return "", fmt.Errorf(
			"copying archive response body for %q: %w",
			archiveName, err,
		)
	}
	archiveHashSum := hex.EncodeToString(archiveHash.Sum(nil))

	checksumsStream, err := gotoolbox.FetchGitHubReleaseAsset(
		repository, version, "checksums.txt",
	)
	if err != nil {
		return "", err
	}
	defer checksumsStream.Close()
	checksumsScanner := bufio.NewScanner(checksumsStream)
	for checksumsScanner.Scan() {
		fields := strings.Fields(checksumsScanner.Text())
		if len(fields) >= 2 && fields[1] == archiveName {
			if fields[0] != archiveHashSum {
				return "", fmt.Errorf(
					"checksum mismatch for %q: expected %q, got %q",
					archiveName, fields[0], archiveHashSum,
				)
			}
			break
		}
	}

	archiveBinName := tool + binExt
	var archiveBinReader io.Reader
	if archiveExt == ".tar.gz" {
		if _, err = archiveTemp.Seek(0, io.SeekStart); err != nil {
			return "", fmt.Errorf(
				"seeking to beginning of archive temp file: %w",
				err,
			)
		}
		gzipReader, err := gzip.NewReader(archiveTemp)
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
					"binary file %q not found in archive %q",
					archiveBinName, archiveName,
				)
			}
			if err != nil {
				return "", fmt.Errorf(
					"reading tar archive %q: %w",
					archiveName, err,
				)
			}
			if header.Typeflag == tar.TypeReg && header.Name == archiveBinName {
				// Found the binary, so exit loop.
				break
			}
		}
		archiveBinReader = tarReader
	} else {
		zipReader, err := zip.NewReader(archiveTemp, archiveSize)
		if err != nil {
			return "", fmt.Errorf("creating zip reader: %w", err)
		}
		archiveBinFile := func() *zip.File {
			for _, file := range zipReader.File {
				if file.Name == archiveBinName {
					return file
				}
			}
			return nil
		}()
		if archiveBinFile == nil {
			return "", fmt.Errorf(
				"binary file %q not found in archive %q",
				archiveBinName, archiveName,
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
	}
	binFile, err := os.OpenFile(
		binPath,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0755,
	)
	if err != nil {
		return "", fmt.Errorf(
			"creating binary file %q: %w",
			binPath, err,
		)
	}
	defer binFile.Close()
	if _, err = io.Copy(binFile, archiveBinReader); err != nil {
		return "", fmt.Errorf(
			"copying archived file to %q: %w",
			binPath, err,
		)
	}

	return binPath, nil
}

func goreleaserPlatform() (string, error) {
	goPlatform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	switch goPlatform {
	case "darwin/amd64":
		return "Darwin_x86_64", nil
	case "darwin/arm64":
		return "Darwin_arm64", nil
	case "linux/386":
		return "Linux_i386", nil
	case "linux/amd64":
		return "Linux_x86_64", nil
	case "linux/arm":
		return "Linux_armv7", nil
	case "linux/arm64":
		return "Linux_arm64", nil
	case "linux/loong64":
		return "Linux_loong64", nil
	case "linux/ppc64":
		return "Linux_ppc64", nil
	case "linux/riscv64":
		return "Linux_riscv64", nil
	case "windows/386":
		return "Windows_i386", nil
	case "windows/amd64":
		return "Windows_x86_64", nil
	case "windows/arm64":
		return "Windows_arm64", nil
	}

	return "", fmt.Errorf("unsupported %s", goPlatform)
}
