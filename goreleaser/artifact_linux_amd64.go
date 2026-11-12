//go:build linux && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Linux_x86_64.tar.gz"
	ArtifactSHA256Digest      = "f3d24969a7389a46613875711d15a96d8f6ef1553d84c6ace09d69ac1e7d44c0"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
