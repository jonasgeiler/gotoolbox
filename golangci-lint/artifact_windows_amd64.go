//go:build windows && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-windows-amd64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-" + Version + "-windows-amd64.zip"
	ArtifactSHA256Digest      = "bd42e3ebc8cb4ececb86941983baaf1dc221bbb04d838e94ce63b49cc91e02bb"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-windows-amd64/golangci-lint.exe"
)
