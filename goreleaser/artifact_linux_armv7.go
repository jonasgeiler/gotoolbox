//go:build linux && arm.7

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-armv7"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.3/goreleaser_Linux_armv7.tar.gz"
	ArtifactSHA256Digest      = "37e6ad1189ad3ef64f9d08468f5c616edb5451156eb95026802e0ed998e1622a"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
