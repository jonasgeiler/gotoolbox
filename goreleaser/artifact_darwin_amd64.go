//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Darwin_x86_64.tar.gz"
	ArtifactSHA256Digest      = "a1dfe9492c67b860129bc08fa5952955f0666abe4741e1846f93a6c6b47360e6"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
