//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.2/goreleaser_Linux_riscv64.tar.gz"
	ArtifactSHA256Digest      = "6f5fce3fca38980cc6290615571b81787a6810a4a940a604306e97d982a758f6"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
