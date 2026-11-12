//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Linux_arm64.tar.gz"
	ArtifactSHA256Digest      = "f0fa989c3531348e8f975aac647eb187a09046f8a34c54e0b79bed66c84b7f96"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
