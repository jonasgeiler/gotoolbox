//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Linux_armv7.tar.gz"
	ArtifactSHA256Digest      = "36c98e07772965d1e79c43395157444417038ea57a1a7f3f54130caa3c90b4c7"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
