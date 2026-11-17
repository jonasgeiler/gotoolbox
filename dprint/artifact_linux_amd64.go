//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "61707303f4305891fecffec1f92ec8243b07e9d3ba8a7e81f7543fc77fb0002d"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "d1aaf0c6d015931b860804eda5bf7fac9d29e2253351b7d6b0177e6b03a8776d"
	ArtifactInArchiveFilePath  = "dprint"
)
