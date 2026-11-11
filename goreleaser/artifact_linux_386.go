//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "5af5ee4bee0348ea08a2df2c231b4fc35c0f75456f50c7b9e057ce699cad0bd0"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
