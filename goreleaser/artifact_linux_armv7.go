//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Linux_armv7.tar.gz"
	ArtifactSHA256Digest      = "ddd2ff2d9f2b3fcd771fcd251e3d00954565a506e7b62333dcd36f9454de1676"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
