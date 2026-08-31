//go:build windows && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-windows-amd64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-windows-amd64.zip"
	ArtifactSHA256Digest      = "4735fdc8e84a0cfb7a15a1c364a650942f88215e0d36c674ebc4024f7b554524"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-windows-amd64/golangci-lint.exe"
)
