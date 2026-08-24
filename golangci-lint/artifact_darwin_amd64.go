//go:build darwin && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v2.12.2-darwin-amd64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-darwin-amd64.tar.gz"
	ArtifactSHA256Digest      = "f6f06d94b6241521c53d15450c5209b028270bf966f842afb11c030c79f5bc16"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-2.12.2-darwin-amd64/golangci-lint"
)
