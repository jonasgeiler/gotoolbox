//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_riscv64.tar.gz"
	ArtifactSHA256Digest      = "d3bfa4f1f4639d45cac045dd129dfabce6ece2b198de0391531b464f67cae273"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
