//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "d2affc275f5c6f6a8e1259842e034942c61b853f8131e51853b7231539283315"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
