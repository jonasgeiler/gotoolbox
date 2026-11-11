//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Windows_i386.zip"
	ArtifactSHA256Digest      = "6594ec1e3410aabd3d20e42a5dc0ad1f51125e259ef69d087db72933c210d1e1"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
