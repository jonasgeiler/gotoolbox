//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-386"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-" + Version + "-linux-386.tar.gz"
	ArtifactSHA256Digest      = "8acadd219d421b89186438c095fd6da72bcb2cc6a334798d31732003c376233a"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-386/golangci-lint"
)
