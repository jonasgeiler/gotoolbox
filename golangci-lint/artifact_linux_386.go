//go:build linux && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-linux-386"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-linux-386.tar.gz"
	ArtifactSHA256Digest      = "2cbf2ff491093fb370f9194e31ef614480135a5ace8239d9f392debfbe9dabb0"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-linux-386/golangci-lint"
)
