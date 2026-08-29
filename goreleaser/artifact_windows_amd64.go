//go:build windows && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-amd64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Windows_x86_64.zip"
	ArtifactSHA256Digest      = "6fe5eda11f0bcac8069aff3ef3dcb0b11816c9e95f89773595564697a5278bc9"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
