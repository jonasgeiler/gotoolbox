//go:build linux && mips64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-mips64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-mips64.tar.gz"
	ArtifactSHA256Digest      = "78cfd916db5c41113133fa8144d4f50c5dfd1445038e56310fbae23bec3c4f89"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-mips64/golangci-lint"
)
