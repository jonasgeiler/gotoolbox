//go:build windows && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Windows_x86_64.zip"
	ArtifactSHA256Digest      = "d230d6def8b6059b277d472a2790501516b429ffe83b91469daad8b7736b8a5d"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
