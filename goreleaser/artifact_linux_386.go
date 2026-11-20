//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.3/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "1b6d5f9b86509c74f554334baa8c615fd1de41c22b2587dddf67b72e763df44c"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
