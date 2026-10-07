//go:build darwin && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-x86_64-apple-darwin.zip"
	ArtifactSHA256Digest       = "9b38610df23417c5cce6daa7d9fb43d9960a5f859fa91cbe7c12f11198074900"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
