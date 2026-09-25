//go:build linux && mips64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-mips64"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-mips64.tar.gz"
	ArtifactSHA256Digest      = "bbb18dedbc66637354d4d76b5fbed9cce0115d76bf9c23895a9b3db1bfc41eca"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-mips64/golangci-lint"
)
