//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "b85ad46a2e8e686b4ff1ff1dc1f0a0f947e766df5afb08ce367cb0d152dff616"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "c33d901d2f5b81af680bce50902e7193e6c169f4f2dd6e1d41d3c5a96d70ecb4"
	ArtifactInArchiveFilePath  = "dprint"
)
