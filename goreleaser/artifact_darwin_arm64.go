//go:build darwin && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.3/goreleaser_Darwin_arm64.tar.gz"
	ArtifactSHA256Digest      = "821c58ab8c31d57bb331cab9f1a703eb3673bff4fce0db9cc3c4a038af14f76f"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
