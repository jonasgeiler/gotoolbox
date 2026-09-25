//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-loong64.tar.gz"
	ArtifactSHA256Digest      = "0dca1c8b3daa6a0eefa152292a413dc8c882ab4d3d7ec4f9efe87ec3e74e837e"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-loong64/golangci-lint"
)
