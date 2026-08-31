//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-loong64.tar.gz"
	ArtifactSHA256Digest      = "15be79776effe89bfbc8c2271026710c24977926ab79d3af01ca80cd63387921"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-loong64/golangci-lint"
)
