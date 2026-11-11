//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Darwin_x86_64.tar.gz"
	ArtifactSHA256Digest      = "77bb1927def4d298c651166c826c0e58e7dc758ec8c1d6893b6c69705719b0f9"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
