//go:build linux && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v2.16.0-linux-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_x86_64.tar.gz"
	ArtifactSHA256Digest      = "eaae05b5eba07533bd0f06846b68c808399504784df00c62eb219541fc04e5e2"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
