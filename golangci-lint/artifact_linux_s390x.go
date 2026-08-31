//go:build linux && s390x

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-s390x"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-s390x.tar.gz"
	ArtifactSHA256Digest      = "5f29782afded451abd5957fb2efdde7f3e1da8b2a0bf395c18be1b7a32070235"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-s390x/golangci-lint"
)
