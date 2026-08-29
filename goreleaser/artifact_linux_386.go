//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.0/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "cfacfeebf9a43d54c74a265af0710674965735daa2c86d97db2e1e456dee8845"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
