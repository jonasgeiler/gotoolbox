package gotoolbox

import (
	"crypto"
	_ "crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type Checksum struct {
	Algorithm crypto.Hash
	Digest    string
}

type Artifact struct {
	CacheName         string
	DownloadURL       string
	Checksum          Checksum
	ArchiveFormat     ArchiveFormat
	InArchiveFilePath string
}

// Tool defines a gotoolbox tool with its download and versioning info.
type Tool struct {
	Name     string
	Artifact Artifact
}

const logPrefix = "gotoolbox | "

func (t *Tool) Run() {
	toolFilePath, err := t.DownloadIfNeeded()
	if err != nil {
		_, printErr := fmt.Fprintf(os.Stderr,
			"%sDownload Error: %v\n",
			logPrefix, err,
		)
		if printErr != nil {
			//goland:noinspection GoErrorStringFormat
			panic(fmt.Errorf("%sDownload Error: %w", logPrefix, err))
		}
		os.Exit(1)
	}

	if err := Exec(toolFilePath); err != nil {
		_, printErr := fmt.Fprintf(os.Stderr,
			"%sExec Error: %v\n",
			logPrefix, err,
		)
		if printErr != nil {
			//goland:noinspection GoErrorStringFormat
			panic(fmt.Errorf("%sExec Error: %w", logPrefix, err))
		}
		os.Exit(1)
	}
}

func (t *Tool) DownloadIfNeeded() (string, error) {
	if t.Artifact.DownloadURL == "" {
		return "", fmt.Errorf(
			"unsupported platform: %s/%s",
			runtime.GOOS, runtime.GOARCH,
		)
	}
	if !t.Artifact.Checksum.Algorithm.Available() {
		return "", fmt.Errorf(
			"unsupported checksum algorithm: %s",
			t.Artifact.Checksum.Algorithm,
		)
	}

	artifactCacheDirPath := filepath.Join(
		CacheDirPath(), // TODO: Use Go workspace/module root instead? Traverse upwards until a go.work file found, if not found try again but look for go.mod. This would also allow just deleting old cached versions maybe.
		t.Artifact.CacheName,
	)

	toolFileName := t.Name
	if runtime.GOOS == "windows" {
		toolFileName += ".exe"
	}
	toolFilePath := filepath.Join(artifactCacheDirPath, toolFileName)
	if _, err := os.Stat(toolFilePath); err == nil {
		// Already exists, skip downloading.
		return toolFilePath, nil
	}

	if err := os.MkdirAll(artifactCacheDirPath, 0755); err != nil {
		return "", fmt.Errorf(
			"creating binary cache directory %q: %w",
			artifactCacheDirPath, err,
		)
	}

	artifactDownload, err := Download(t.Artifact.DownloadURL)
	if err != nil {
		return "", fmt.Errorf(
			"downloading file from %q: %w",
			t.Artifact.DownloadURL, err,
		)
	}
	defer artifactDownload.Close()

	artifactDownloadFile, err := os.CreateTemp(
		artifactCacheDirPath,
		".temp-download-*",
	)
	if err != nil {
		return "", fmt.Errorf(
			"creating temporary file %q: %w",
			artifactCacheDirPath, err,
		)
	}
	defer os.Remove(artifactDownloadFile.Name())
	defer artifactDownloadFile.Close()

	artifactDownloadHash := t.Artifact.Checksum.Algorithm.New()
	artifactDownloadFileSize, err := io.Copy(
		io.MultiWriter(artifactDownloadFile, artifactDownloadHash),
		artifactDownload,
	)
	if err != nil {
		return "", fmt.Errorf(
			"streaming/copying response body from %q: %w",
			t.Artifact.DownloadURL, err,
		)
	}
	artifactDownloadDigest := hex.EncodeToString(artifactDownloadHash.Sum(nil))

	if artifactDownloadDigest != t.Artifact.Checksum.Digest {
		return "", fmt.Errorf(
			"checksum mismatch for file downloaded from %q: expected %q, got %q",
			t.Artifact.DownloadURL,
			t.Artifact.Checksum.Digest,
			artifactDownloadDigest,
		)
	}

	var tempToolFile *os.File
	if t.Artifact.ArchiveFormat == NotAnArchive {
		// We have directly downloaded a binary and can just re-use it.
		tempToolFile = artifactDownloadFile
	} else {
		// We have downloaded an archive and need to extract it.
		inArchiveToolFile, err := OpenFileInArchive(
			artifactDownloadFile,
			artifactDownloadFileSize,
			t.Artifact.ArchiveFormat,
			t.Artifact.InArchiveFilePath,
		)
		if err != nil {
			return "", err
		}
		defer inArchiveToolFile.Close()

		tempToolFile, err = os.CreateTemp(
			artifactCacheDirPath,
			".temp-tool-*",
		)
		if err != nil {
			return "", fmt.Errorf(
				"creating temporary file in %q: %w",
				artifactCacheDirPath, err,
			)
		}
		defer os.Remove(tempToolFile.Name())
		defer tempToolFile.Close()

		if _, err := io.Copy(
			tempToolFile,
			inArchiveToolFile,
		); err != nil {
			return "", fmt.Errorf(
				"extracting file from %q to %q: %w",
				artifactDownloadFile.Name(), tempToolFile.Name(), err,
			)
		}
	}

	if err := tempToolFile.Chmod(0755); err != nil {
		return "", fmt.Errorf(
			"changing file permissions for %q: %w",
			tempToolFile.Name(), err,
		)
	}
	if err := os.Rename(tempToolFile.Name(), toolFilePath); err != nil {
		return "", fmt.Errorf(
			"moving temporary file %q to %q: %w",
			tempToolFile.Name(), toolFilePath, err,
		)
	}

	return toolFilePath, nil
}
