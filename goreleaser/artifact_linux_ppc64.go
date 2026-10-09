//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.3/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "cd11abc60244b7cd0d224679bed8ffb908384cde0d82562a4694de16ea285dc9"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
