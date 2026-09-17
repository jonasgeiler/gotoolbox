//go:build linux && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Linux_x86_64.tar.gz"
	ArtifactSHA256Digest      = "0a96edc9d9bc594e4a41cc4d59467c182062910ab24d9d1f6dd7b667d32606d3"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
