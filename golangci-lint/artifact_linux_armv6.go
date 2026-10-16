//go:build linux && arm.6

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-armv6"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.15.0/golangci-lint-" + Version + "-linux-armv6.tar.gz"
	ArtifactSHA256Digest      = "2157ff45d57a83f61f329af01807ae828c35768be4d022e8acbb63b66e820f2e"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-armv6/golangci-lint"
)
