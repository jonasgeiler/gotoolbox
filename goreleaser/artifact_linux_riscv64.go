//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.0/goreleaser_Linux_riscv64.tar.gz"
	ArtifactSHA256Digest      = "24be27a0e74a51bb5dc5c09dd03dc959d5cdf742329d60095b90d9921541d509"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
