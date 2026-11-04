//go:build darwin && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Darwin_arm64.tar.gz"
	ArtifactSHA256Digest      = "50bd83a6fcbf1dac272cfc4c46f9116052a9b0b463404a044af71a832fd92ad6"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
