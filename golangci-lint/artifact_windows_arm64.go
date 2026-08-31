//go:build windows && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-windows-arm64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-windows-arm64.zip"
	ArtifactSHA256Digest      = "2dbffbd1225d41ac5740f0b478a43b6517f3e3f702fe0ab3aec470bd6ec8e263"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-windows-arm64/golangci-lint.exe"
)
