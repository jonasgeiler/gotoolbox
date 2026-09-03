//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "371b63109cbd7b34e179fc3af0815f65a5d9f3a560d2a245a78ef99ec71b3f9e"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "a73695f6407c6c36adef5971053f81ccef13b0a04b6b1da98312ae0ee5332edd"
	ArtifactInArchiveFilePath  = "dprint"
)
