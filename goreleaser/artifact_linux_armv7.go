//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.19.0/goreleaser_Linux_armv7.tar.gz"
	ArtifactSHA256Digest      = "9b5f40da156e6e2c56fa3dd9634737400198427a00fcf07873d97e97836c5850"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
