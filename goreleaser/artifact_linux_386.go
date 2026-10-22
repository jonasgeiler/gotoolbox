//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.19.0/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "b889f692f4a713e32afcea2edf9c201298155648b89fc0826c746776f28ff9cc"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
