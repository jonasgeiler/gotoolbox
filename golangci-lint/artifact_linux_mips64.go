//go:build linux && mips64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-linux-mips64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-mips64.tar.gz"
	ArtifactSHA256Digest      = "d73c73e3f3090659e6ec1276e1f8497d9084690476d7d004672dae1199550b1c"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-linux-mips64/golangci-lint"
)
