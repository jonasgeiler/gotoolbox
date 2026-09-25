//go:build linux && arm.6

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-armv6"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-armv6.tar.gz"
	ArtifactSHA256Digest      = "20d75528b7b7fa37dd83101b0be8b349e4f65f6e970d5fea0a96c7959e6e375a"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-armv6/golangci-lint"
)
