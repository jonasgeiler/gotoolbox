//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Linux_loong64.tar.gz"
	ArtifactSHA256Digest      = "c0e8492973d2c20f0dbcb654ddab0d6a6747fb4d9aea3e176916f5419a57a801"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
