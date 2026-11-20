//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.3/goreleaser_Linux_loong64.tar.gz"
	ArtifactSHA256Digest      = "05d527a7d524879ac219b44c8a5005e11fcd8d685a471812fbb30ca8dbd56972"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
