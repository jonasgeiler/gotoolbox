//go:build darwin && arm64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-darwin-arm64.tar.gz"
	ArtifactSHA256Digest      = "f4bf83f0b64f055c42b28fc9a38861839f69c096e61c788e72dfaae412011789"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-darwin-arm64/golangci-lint"
)
