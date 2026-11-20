//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.3/goreleaser_Linux_arm64.tar.gz"
	ArtifactSHA256Digest      = "cfce0afc7bcc14447df211cc10c0012cfc395646070bcf23e840b16cbce6d204"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
