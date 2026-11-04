//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.20.0/goreleaser_Linux_riscv64.tar.gz"
	ArtifactSHA256Digest      = "24f0a36322deda5f4c4668f336013f98d9652393c10ee5c38c2745da59525622"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
