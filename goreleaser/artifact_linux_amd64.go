//go:build linux && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Linux_x86_64.tar.gz"
	ArtifactSHA256Digest      = "79e0589ca9efd70e8db85158a21334cbb57811d4034b39b61d55789f361341a5"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
