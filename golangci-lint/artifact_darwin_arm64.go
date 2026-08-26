//go:build darwin && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-darwin-arm64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-darwin-arm64.tar.gz"
	ArtifactSHA256Digest      = "a9c54498731b3128f79e090be6110f3e5fffccc617b08142ed244d4126c73f29"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-darwin-arm64/golangci-lint"
)
