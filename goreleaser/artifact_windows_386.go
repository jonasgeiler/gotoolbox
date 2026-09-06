//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.1/goreleaser_Windows_i386.zip"
	ArtifactSHA256Digest      = "bf9c7a9e0dd53c4b211859b0c7fd9783efbecb1cd6f5de8b3bf3bd022f4a3c39"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
