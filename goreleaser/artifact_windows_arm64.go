//go:build windows && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Windows_arm64.zip"
	ArtifactSHA256Digest      = "961784dac3a74c9bb58c3df61f33163230df5c8dfcf9af68c84a838358a04ebc"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
