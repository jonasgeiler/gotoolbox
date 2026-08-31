//go:build linux && arm.6

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-armv6"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-armv6.tar.gz"
	ArtifactSHA256Digest      = "1d5100291e3eed1a4e55dc094ebdac9926ce66def857dd7ef80504241695f3b1"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-armv6/golangci-lint"
)
