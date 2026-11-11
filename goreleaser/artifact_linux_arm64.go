//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Linux_arm64.tar.gz"
	ArtifactSHA256Digest      = "81dc6c218288358e8c0b021ee1fceb6a6a3ec4f29b11bf830cab985507af492d"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
