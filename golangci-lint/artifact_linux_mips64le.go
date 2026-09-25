//go:build linux && mips64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-mips64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-mips64le.tar.gz"
	ArtifactSHA256Digest      = "077c1bf1421ec29c7f28b5e9ebd7692f8521dcc2ca7baa34b57b900f65c15e43"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-mips64le/golangci-lint"
)
