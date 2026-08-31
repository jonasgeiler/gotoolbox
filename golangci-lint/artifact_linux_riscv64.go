//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-riscv64.tar.gz"
	ArtifactSHA256Digest      = "170f60a7228e3c1eda8e56041fffd11c6aba9b47f820005a04f381aa31871316"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-riscv64/golangci-lint"
)
