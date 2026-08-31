//go:build windows && 386

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "golangci-lint-v" + Version + "-windows-386"
	ArtifactDownloadURL       = "https://github.com/golangci/golangci-lint/releases/download/v2.13.2/golangci-lint-" + Version + "-windows-386.zip"
	ArtifactSHA256Digest      = "15cad8e2347d91b259095c29c589b53baaf051893a7b4b44861cad09a3069e2e"
	ArtifactArchiveFormat     = gotoolbox.ZipArchive
	ArtifactInArchiveFilePath = "golangci-lint-" + Version + "-windows-386/golangci-lint.exe"
)
