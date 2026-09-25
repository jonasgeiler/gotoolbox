//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-darwin-amd64.tar.gz"
	ArtifactSHA256Digest      = "a5667c1c3536be1740133213e1e822bfb8f0d98ea12903174d6d5f635e4ed68d"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-darwin-amd64/golangci-lint"
)
