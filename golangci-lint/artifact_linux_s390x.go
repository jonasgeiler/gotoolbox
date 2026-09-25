//go:build linux && s390x

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-s390x"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-s390x.tar.gz"
	ArtifactSHA256Digest      = "788d7f54c9d141be5942e9c9ca8a56bbea44e5c4dbc2e774e54278bd930e8379"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-s390x/golangci-lint"
)
