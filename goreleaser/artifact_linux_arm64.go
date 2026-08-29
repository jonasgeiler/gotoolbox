//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.0/goreleaser_Linux_arm64.tar.gz"
	ArtifactSHA256Digest      = "1975566c9668e6f4247e6bb57656f21da13635c24d948ef47b1232e5c864a35b"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
