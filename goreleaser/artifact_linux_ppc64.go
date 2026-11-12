//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "40e6d80db621dc92c3422b227fcceb679af5ba8ae094599abbec2d4a7ce2ed2a"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
