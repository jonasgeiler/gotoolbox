//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.1/goreleaser_Linux_loong64.tar.gz"
	ArtifactSHA256Digest      = "906e7f280c256678dfe0b52ee3e76ec5e869d6bf8e08dae93bc5e043cd053d3d"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
