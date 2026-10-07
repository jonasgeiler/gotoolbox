//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "d0658fcc2887fe8ba396f6ef0fec28c97a6d59558835b7b82e573ea923c53b25"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "6585c1a475e995cbd3729e039579c6fef9f28202c3b94bf252ab04f8f4177876"
	ArtifactInArchiveFilePath  = "dprint"
)
