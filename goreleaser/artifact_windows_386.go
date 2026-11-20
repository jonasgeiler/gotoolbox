//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.3/goreleaser_Windows_i386.zip"
	ArtifactSHA256Digest      = "43d98e7dfdc4ca572f0264167ad4e2304d834d3869c7b48e1d7897f11fd4d499"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
