//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "c8826c0b72fc0e1655404bf91fc29f821c01436e53152788e417df13a40d0671"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "5d83774c7546bd7a223f243a17e7851a75b8315899d2a42268e40ebbae66c69b"
	ArtifactInArchiveFilePath  = "dprint"
)
