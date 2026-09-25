//go:build linux && riscv64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-riscv64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-riscv64.tar.gz"
	ArtifactSHA256Digest      = "c36e17e0bd02dd00c0db5b64ba21afb58652e5b7aeca17479dcba98906d8d7c1"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-riscv64/golangci-lint"
)
