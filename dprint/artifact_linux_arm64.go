//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "e1b713806f7f7b94072ba2b069aa19ac28f475dbb1fbc5b672d3030a2ab001ee"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "ff961314e7c1ca6a02eaada0e37920a37c668fbc09fd85838c2aced41a4bc0e4"
	ArtifactInArchiveFilePath  = "dprint"
)
