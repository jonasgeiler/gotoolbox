//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "c74dd4f48dd9b8d6d595acfb238dd32ede67bdf792e8c6e2b8760cdaff9d80fb"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "f161161399b5ef62b3b68570ea1df24d16687adfbb88c125c8c4355e5780c407"
	ArtifactInArchiveFilePath  = "dprint"
)
