//go:build linux && mips64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-mips64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-" + Version + "-linux-mips64le.tar.gz"
	ArtifactSHA256Digest      = "55429607fb7608f3b1748ece9ab4a74a3eec46ffcaca114bfaf6c0f3d70d4e0d"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-mips64le/golangci-lint"
)
