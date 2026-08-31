//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-arm64.tar.gz"
	ArtifactSHA256Digest      = "a2a4e0065aa41be71f7c5ac90f271b61751331e5d04314e62afe4027855f0893"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-arm64/golangci-lint"
)
