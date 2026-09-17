//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Windows_i386.zip"
	ArtifactSHA256Digest      = "6de06d4ad1200483eff45f51f0274131a9b1b2f1a057ec9b8aba2825120a8557"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
