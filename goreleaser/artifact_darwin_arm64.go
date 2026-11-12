//go:build darwin && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Darwin_arm64.tar.gz"
	ArtifactSHA256Digest      = "4e68482f6bbdb2d3ce7f835be50a4eb7287ee999cefc563c49c5f42ad3778c78"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
