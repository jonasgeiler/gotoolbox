//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.1/goreleaser_Linux_riscv64.tar.gz"
	ArtifactSHA256Digest      = "5306dc26437892a5acfc6c3ecc92b43a0a3dfab55156ce4c8473a5cf17da275c"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
