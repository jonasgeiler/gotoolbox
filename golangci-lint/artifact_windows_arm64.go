//go:build windows && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-windows-arm64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-windows-arm64.zip"
	ArtifactSHA256Digest      = "947b9a5bf762d465710b376c156f0184abb2168378b0826af1899e0ee7183742"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-windows-arm64/golangci-lint.exe"
)
