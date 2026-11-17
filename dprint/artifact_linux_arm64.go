//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "a427ef844b300d896b3f836706bb77d8df6b77f8f2627103c4a34ded94cf6ae1"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "804535e54f431674f5481725ba5d4718de0a7398d6abd33b899fd734f6d98062"
	ArtifactInArchiveFilePath  = "dprint"
)
