//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "c1e709d629caf8c1f8ef8328f2b5136c68dd5e878b36b04403bf9dfd133c559f"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "1705b1be1773955e882a69c29aebd3862d5b466c951f47495ad6cb423c224f86"
	ArtifactInArchiveFilePath  = "dprint"
)
