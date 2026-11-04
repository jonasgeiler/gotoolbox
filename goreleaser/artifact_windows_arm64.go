//go:build windows && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-arm64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Windows_arm64.zip"
	ArtifactSHA256Digest      = "b7ae30d558db74d67bd03e39d3309aabf47e1a18649ef4ac4429f329c357f496"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
