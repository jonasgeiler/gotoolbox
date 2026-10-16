//go:build linux && mips64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-mips64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.15.0/golangci-lint-" + Version + "-linux-mips64le.tar.gz"
	ArtifactSHA256Digest      = "fa07b5795eeeb7a80eb99bb96c582882b628c01cebde7ad4d9c3ee1f7395f47d"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-mips64le/golangci-lint"
)
