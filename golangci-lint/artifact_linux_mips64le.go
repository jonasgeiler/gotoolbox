//go:build linux && mips64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-mips64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-mips64le.tar.gz"
	ArtifactSHA256Digest      = "fdf427ff79f131a3a6874bfff4e6cfa569e9fbd66bd360a0d7cb206c1d9f5bff"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-mips64le/golangci-lint"
)
