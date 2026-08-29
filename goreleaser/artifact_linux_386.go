//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "b6b0764b9e339fcfd8bcf1786424c99ddecbecf27d15025c189e1c64932a1563"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
