//go:build linux && ppc64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-ppc64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-" + Version + "-linux-ppc64le.tar.gz"
	ArtifactSHA256Digest      = "31561f2e35ca8e2b9f8c2bc3055c74dd3f0fd341db7c9d0feb5292c95bda1a98"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-ppc64le/golangci-lint"
)
