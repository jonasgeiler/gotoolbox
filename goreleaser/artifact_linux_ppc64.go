//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "f3fcbd2b54a2c699de753787e8550bd91c82d3a705df9bf611fa2cfe376635d0"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
