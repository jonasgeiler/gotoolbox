//go:build darwin && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Darwin_arm64.tar.gz"
	ArtifactSHA256Digest      = "07a98d41abe1678020a8245d98acdc6b9d30f850471fe40f90eeaf57af74046e"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
