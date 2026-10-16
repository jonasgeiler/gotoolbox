//go:build linux && s390x

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-s390x"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.15.0/golangci-lint-" + Version + "-linux-s390x.tar.gz"
	ArtifactSHA256Digest      = "8571d506e9d1a076876bef95fb5ef99b562216782f710ff81a7130d0870cfb6e"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-s390x/golangci-lint"
)
