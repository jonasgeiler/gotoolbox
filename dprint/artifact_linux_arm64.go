//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "b90a0e7ede5fbae3e9f4e0e50455902a7e436f575c8db3b9b7fea9b7a0303210"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "4c767249bd35be5d3a0fab98a2eb94bcc439330a45d40c2681032d27ed1032b1"
	ArtifactInArchiveFilePath  = "dprint"
)
