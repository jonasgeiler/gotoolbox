//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.19.0/goreleaser_Windows_i386.zip"
	ArtifactSHA256Digest      = "796e2d925fad9c79cea940265d2503dc8328401a125b8f4a8fffef0c1a0f4c75"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
