//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-windows-386"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-windows-386.zip"
	ArtifactSHA256Digest      = "6242506521a9fba4ba3d86f7d2842d284dcd144ca0f95671ce52c6b0b22a6417"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-windows-386/golangci-lint.exe"
)
