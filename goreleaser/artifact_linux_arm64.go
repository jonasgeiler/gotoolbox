//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Linux_arm64.tar.gz"
	ArtifactSHA256Digest      = "2dada0a6bbcd10f3a05228d98e77d25fb362efb9e5bab6362cf37e64c8a26cb9"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
