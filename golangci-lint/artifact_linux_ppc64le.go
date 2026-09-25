//go:build linux && ppc64le

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-ppc64le"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.14.0/golangci-lint-" + Version + "-linux-ppc64le.tar.gz"
	ArtifactSHA256Digest      = "5189f8dbe1afb8fb394a9c3c29513c2e3a38790405146397e2ab08e243a133cf"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-ppc64le/golangci-lint"
)
