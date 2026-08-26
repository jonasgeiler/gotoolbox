//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-riscv64.tar.gz"
	ArtifactSHA256Digest      = "32d67a82e5711519aa44ec415e0cb6d1fad9e8d390a95c81e9aeeb1e8a1bf211"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-linux-riscv64/golangci-lint"
)
