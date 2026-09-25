//go:build windows && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-windows-arm64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-windows-arm64.zip"
	ArtifactSHA256Digest      = "6399496aef883e5d2f77ea93e6e54fdd05bc2695f64071d8de0b0579cad92274"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-windows-arm64/golangci-lint.exe"
)
