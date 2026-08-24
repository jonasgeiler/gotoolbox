//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-arm64.tar.gz"
	ArtifactSHA256Digest      = "44cd40a8c76c86755375adfeea52cfd3533cb43d7bd647771e0ae065e166df3a"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-linux-arm64/golangci-lint"
)
