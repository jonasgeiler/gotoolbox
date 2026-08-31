//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "4458fb5ad357c2735210470a91e3f2a95abafcd7f2d7840cb94afc96dd11fbfe"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "76f01495dae94d531e9ebd5e97977ae9a8edc8d7fc8ae8a479d34715617cb69d"
	ArtifactInArchiveFilePath  = "dprint"
)
