//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Windows_i386.zip"
	ArtifactSHA256Digest      = "a0a72653b99eefd9855f69b9bbb717757956881f4c2190416a5a1c505eeb9de4"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
