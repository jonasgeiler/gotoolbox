//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-darwin-amd64.tar.gz"
	ArtifactSHA256Digest      = "8a13aaf9cbbb1dee52824e862cf0d0720e5bb97c1f4260d1e51623a09492b57b"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-darwin-amd64/golangci-lint"
)
