//go:build darwin && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-aarch64-apple-darwin.zip"
	ArtifactSHA256Digest       = "d2a6e67e6aa308ed13d09627e8a8e8fe2cd3f62aeb406fe653d59bcec2c25c47"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
