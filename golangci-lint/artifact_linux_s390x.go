//go:build linux && s390x

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-s390x"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-" + Version + "-linux-s390x.tar.gz"
	ArtifactSHA256Digest      = "e4a35d5531c8c3967e6651f763e371540d2d736972161c2cea35c7601ac0168c"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-s390x/golangci-lint"
)
