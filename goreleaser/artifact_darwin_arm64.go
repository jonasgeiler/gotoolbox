//go:build darwin && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.3/goreleaser_Darwin_arm64.tar.gz"
	ArtifactSHA256Digest      = "c1423474ab35ab6d7669d87d403dc9ddebf6013a11d3de7309ba840765a19af1"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
