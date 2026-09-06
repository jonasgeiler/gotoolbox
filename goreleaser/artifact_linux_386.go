//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.1/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "9c9f2649b659e65c2cb3ad08445df0be5eab51ffa5aa3a09a797aaa8b7beefd8"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
