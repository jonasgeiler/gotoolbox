//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "c423c7ee51187129db67f0bab4584cf7ff6fdeb29faf04d1521fa79a5c12c520"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "79025432b0fcfaa09fa1324a97a2cf4ec5d972175830fb496dc4be4f6a89340e"
	ArtifactInArchiveFilePath  = "dprint"
)
