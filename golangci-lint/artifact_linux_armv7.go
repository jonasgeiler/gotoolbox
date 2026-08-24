//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-armv7.tar.gz"
	ArtifactSHA256Digest      = "40602c69b04f5262aac21ce090aafb560c4299eadd31dbdc158c074cc4cf9789"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-linux-armv7/golangci-lint"
)
