//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName       = "goreleaser-v" + Version + "-linux-386"
	ArtifactDownloadURL     = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.3/goreleaser_Linux_i386.tar.gz"
	ArtifactSHA256Digest    = "4fbc575c3d433ec7f746d1cbae33ab7f46db8073e60c10d5d67171d01b1fb960"
	ArtifactArchinAveFormat = gotoolbox.TarGzipArchive
	ArtifactToolIrchive     = "goreleaser"
)
