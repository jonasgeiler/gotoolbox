//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "bbed47893c417e5452caf857c234d71617a45cd7102bb743078966a0485ea5ce"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
