//go:build windows && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v2.16.0-windows-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Windows_arm64.zip"
	ArtifactSHA256Digest      = "1183c81863044ce9baa89c1393c258949390b8df683df7ca959e9c718d7723c9"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
