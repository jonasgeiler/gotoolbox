//go:build windows && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Windows_x86_64.zip"
	ArtifactSHA256Digest      = "b84ce9966dc8313edb935846431123b8a9d249eb7140a44783a1a45a5addcf5d"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
