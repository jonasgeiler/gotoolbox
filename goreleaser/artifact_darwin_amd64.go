//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.0/goreleaser_Darwin_x86_64.tar.gz"
	ArtifactSHA256Digest      = "c115f9ca07163d55885ba2276c5c2efebc95d60f7f7f69fe2dd6a54e97ac6db4"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
