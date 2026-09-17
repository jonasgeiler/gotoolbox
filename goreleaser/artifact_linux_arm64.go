//go:build linux && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Linux_arm64.tar.gz"
	ArtifactSHA256Digest      = "a71681b29194f08f057a68cfcaa5c6b15d907a83a2622c51900c4faff828f322"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
