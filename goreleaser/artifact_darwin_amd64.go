//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.3/goreleaser_Darwin_x86_64.tar.gz"
	ArtifactSHA256Digest      = "dbc39347d9ca25a2da9bcc03633f39130bb3a67c9f2a0da86ba15b8b00bf86ce"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
