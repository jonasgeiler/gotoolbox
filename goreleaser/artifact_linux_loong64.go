//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.19.0/goreleaser_Linux_loong64.tar.gz"
	ArtifactSHA256Digest      = "cee780eb4eb58baf76a465e93f71d3963268e946bef24cb7974d15ef928f2ec8"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
