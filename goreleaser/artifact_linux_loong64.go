//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Linux_loong64.tar.gz"
	ArtifactSHA256Digest      = "c407e98d0d7b4dfa371372cfc0a6098e49c811b457fb684d7f7991c6842e8edc"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
