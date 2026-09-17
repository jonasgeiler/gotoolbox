//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Darwin_x86_64.tar.gz"
	ArtifactSHA256Digest      = "5e97d6517f73a0b6f71675a2911b15d3136c03c8907650c2b18e114b1c0f5205"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
