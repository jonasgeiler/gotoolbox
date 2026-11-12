//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "243fa9a298bfe893d7c3d9fb25c1eb819fc57b3968fe5364c727aef9b1afcfb6"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
