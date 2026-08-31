//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "6f9541637ff47289409d709e4b006eeed84fd11c212dc029c4d26504a2cae0ff"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "e9dc29baca00edf30d66b1b7a8de490c3a1bda4091bbc7b69f01f4a88db06c01"
	ArtifactInArchiveFilePath  = "dprint"
)
