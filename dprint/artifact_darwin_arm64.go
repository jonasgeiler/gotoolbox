//go:build darwin && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-aarch64-apple-darwin.zip"
	ArtifactSHA256Digest       = "14bdbfb9b1e3ebd614d5ec08a068b04b0d0370d70c112723003dd9103a4c0c3a"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
