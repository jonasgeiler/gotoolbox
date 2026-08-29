//go:build linux && arm.6

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-armv6"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-" + Version + "-linux-armv6.tar.gz"
	ArtifactSHA256Digest      = "871f97d1a6a8dd8eb2153ec8e1addfc0d2633f42dac1cc8461508a23f971e99d"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-armv6/golangci-lint"
)
