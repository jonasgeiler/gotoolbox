package main

import (
	"archive/zip"
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/jonasgeiler/gotoolbox"
)

const tool = "dprint"

const (
	// renovate: datasource=github-releases depName=dprint/dprint
	version    = "0.54.0"
	repository = "dprint/dprint"
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
	platform, err := dprintPlatform()
	if err != nil {
		return "", fmt.Errorf(
			"determining %s platform: %w",
			tool, err,
		)
	}
	binExt := ""
	if runtime.GOOS == "windows" {
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

	archiveName := fmt.Sprintf("dprint-%s.zip", platform)
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
		repository, version, "SHASUMS256.txt",
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

	zipReader, err := zip.NewReader(archiveTemp, archiveSize)
	if err != nil {
		return "", fmt.Errorf("creating zip reader: %w", err)
	}
	archiveBinName := tool + binExt
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
	archiveBinReader, err := archiveBinFile.Open()
	if err != nil {
		return "", fmt.Errorf(
			"opening file %q in zip archive: %w",
			archiveBinFile.Name, err,
		)
	}
	defer archiveBinReader.Close()
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

func dprintPlatform() (string, error) {
	goPlatform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	switch goPlatform {
	case "darwin/amd64":
		return "x86_64-apple-darwin", nil
	case "darwin/arm64":
		return "aarch64-apple-darwin", nil
	case "linux/amd64":
		if gotoolbox.IsMuslLibc() {
			return "x86_64-unknown-linux-musl", nil
		}
		return "x86_64-unknown-linux-gnu", nil
	case "linux/arm64":
		if gotoolbox.IsMuslLibc() {
			return "aarch64-unknown-linux-musl", nil
		}
		return "aarch64-unknown-linux-gnu", nil
	case "linux/loong64":
		if gotoolbox.IsMuslLibc() {
			return "loongarch64-unknown-linux-musl", nil
		}
		return "loongarch64-unknown-linux-gnu", nil
	case "linux/riscv64":
		if gotoolbox.IsMuslLibc() {
			return "", fmt.Errorf("unsupported %s (with musl libc)", goPlatform)
		}
		return "riscv64gc-unknown-linux-gnu", nil
	case "windows/amd64":
		return "x86_64-pc-windows-msvc", nil
	}

	return "", fmt.Errorf("unsupported %s", goPlatform)
}
