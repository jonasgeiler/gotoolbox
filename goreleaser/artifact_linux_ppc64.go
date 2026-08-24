//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v2.16.0-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "34920ed822616e10216069fec380c832ffad0501d7a9de6680aa103169e940b6"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
