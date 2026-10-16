//go:build linux && mips64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-mips64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.15.0/golangci-lint-" + Version + "-linux-mips64.tar.gz"
	ArtifactSHA256Digest      = "9018d14cd07d46d1b8f933e988ef9ae964e36a2ce19a6138eb10c7a0e237a516"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-mips64/golangci-lint"
)
