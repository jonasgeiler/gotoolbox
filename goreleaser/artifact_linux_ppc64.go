//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "5baaae2ca9d013fa19b01f28b87a7636cb145324f5102e87b70cadc3017059ae"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
