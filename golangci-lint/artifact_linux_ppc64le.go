//go:build linux && ppc64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-ppc64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-ppc64le.tar.gz"
	ArtifactSHA256Digest      = "a1e1a2610289fedaa2c787450819607a047537219fe65980f1bd0aa7ef71c53b"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-ppc64le/golangci-lint"
)
