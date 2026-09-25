//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-armv7.tar.gz"
	ArtifactSHA256Digest      = "2b5a9ef85855cdc21506e0b0e08d0f9cfab0bb1339b61bba337e864a40da0518"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-armv7/golangci-lint"
)
