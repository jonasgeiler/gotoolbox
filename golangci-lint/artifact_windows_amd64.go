//go:build windows && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-windows-amd64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-windows-amd64.zip"
	ArtifactSHA256Digest      = "b15f903d1649283a2ab6e8d2f0c317eaf1f9f6b3763950e7e8b7e72a43dd724f"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-windows-amd64/golangci-lint.exe"
)
