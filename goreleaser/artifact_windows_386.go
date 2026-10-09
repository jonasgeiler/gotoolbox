//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.3/goreleaser_Windows_i386.zip"
	ArtifactSHA256Digest      = "2ec7663916150fe8b4775fa2fb0911c7e3fd0de80c8c28f0efa6abbbbb2a3ba2"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "goreleaser.exe"
)
