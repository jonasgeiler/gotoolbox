//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-windows-386.zip"
	ArtifactSHA256Digest      = "deef62dbace0236fa67f449b3422ce4cefc82d815c2abe5b66e85a8fbf8aca84"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-windows-386/golangci-lint.exe"
)
