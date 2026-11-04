//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "9a41a13d6cd6f51d4fdae6f5685c7989761abe4c444e69162899ecb16e983d26"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
