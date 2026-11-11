//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Linux_loong64.tar.gz"
	ArtifactSHA256Digest      = "0ec2735a9412b60224633617aecce719b72fcb9a0fa3177c01b496fd25002b8c"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
