//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Linux_armv7.tar.gz"
	ArtifactSHA256Digest      = "ec0948ac04cf0197241ecb03bc3f003a6bbabfa8d12908d7cfc3a8ee0803c67f"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
