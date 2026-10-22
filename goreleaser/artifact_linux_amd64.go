//go:build linux && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.19.0/goreleaser_Linux_x86_64.tar.gz"
	ArtifactSHA256Digest      = "ebc3194ad1d1b1787614405285a0039665cd4c5bd3305215a1b201600f90630b"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
