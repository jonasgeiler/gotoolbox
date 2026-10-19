//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "6366cda97e81de1bbc54d91d461e857afa595feb568bd1bec94ab2122e13d969"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "f9308f92d90ad8c5487730b5d8f90757e0e51070f075113d0e654f97f5d72a45"
	ArtifactInArchiveFilePath  = "dprint"
)
