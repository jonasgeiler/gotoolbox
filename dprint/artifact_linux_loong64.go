//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "76aa022f9722f94222240cb84b44df9e8e64d9b4a4a6acde01f8aa9b7cc4e776"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "3c7771c4026c2769c0d77675caa4a426d025e7403e6c911d51a3f36385c8f1a4"
	ArtifactInArchiveFilePath  = "dprint"
)
