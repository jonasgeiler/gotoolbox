//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "7ff2c5c9fcbf0b7abf189557bd8538c487083eb304ab16f444c439ee6b307a83"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "9af469e2b6e0af97944675c611b4d6b23902d8ee2000049f920fa928fb230d16"
	ArtifactInArchiveFilePath  = "dprint"
)
