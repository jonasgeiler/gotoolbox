//go:build linux && amd64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-amd64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-amd64.tar.gz"
	ArtifactSHA256Digest      = "2277d43b98ec0054280f2ac26b53268bae97682444678a59a657dd565da021d6"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-amd64/golangci-lint"
)
