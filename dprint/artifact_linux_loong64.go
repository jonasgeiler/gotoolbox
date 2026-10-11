//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "0dc2f45f9ab2de05f55fad978bda72e2f4b6befda7ada60f52998a0ba002ecbd"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "8e84daaa12f331cceb51bcf45965c4a012ad038b780ccaff5f99b58c131f30e0"
	ArtifactInArchiveFilePath  = "dprint"
)
