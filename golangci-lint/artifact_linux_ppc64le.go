//go:build linux && ppc64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-ppc64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.15.0/golangci-lint-" + Version + "-linux-ppc64le.tar.gz"
	ArtifactSHA256Digest      = "14df8d543f92937ddf2889d3c1fa44277e6c774c6a75354ecc81cf1204b4b7b8"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-ppc64le/golangci-lint"
)
