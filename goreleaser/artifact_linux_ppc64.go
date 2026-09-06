//go:build linux && ppc64

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = "goreleaser-v" + Version + "-linux-ppc64"
	ArtifactDownloadURL       = "https://github.com/goreleaser/goreleaser/releases/download/v2.18.1/goreleaser_Linux_ppc64.tar.gz"
	ArtifactSHA256Digest      = "bbf2a31e3344a29f3ff1a889e4868c7d68f244a70dd9604671bf211b0d6c0ede"
	ArtifactArchiveFormat     = gotoolbox.TarGzipArchive
	ArtifactInArchiveFilePath = "goreleaser"
)
