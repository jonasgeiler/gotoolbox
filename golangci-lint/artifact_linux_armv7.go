//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-armv7.tar.gz"
	ArtifactSHA256Digest      = "01648f24c70b37a6d2e240a2696d35e6ca3fe5bd1215815624e7cac90e6071f7"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-armv7/golangci-lint"
)
