//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.15.0/golangci-lint-" + Version + "-linux-armv7.tar.gz"
	ArtifactSHA256Digest      = "4b3b298b9055447b34f28dfac6d19a6c42d3a6f2a943af93d38a63ee96be2858"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-armv7/golangci-lint"
)
