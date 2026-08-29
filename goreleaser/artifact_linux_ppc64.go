//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.0/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "c5953c1b7655d7b0336c7731a30c124842841c4e8f6f765fb1e9e00701a634cc"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
