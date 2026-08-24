//go:build linux && loong64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-linux-loong64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-loong64.tar.gz"
	ArtifactSHA256Digest      = "76bfc32dff3597190d1409621c18baa31698c87c52c5b8a7c3c86fdb540c4d73"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-linux-loong64/golangci-lint"
)
