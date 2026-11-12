//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Linux_riscv64.tar.gz"
	ArtifactSHA256Digest      = "72928916d670beef71e2e6232627f511abfac08a47673fead8800f8a5c70e84f"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
