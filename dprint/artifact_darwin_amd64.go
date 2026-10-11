//go:build darwin && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-x86_64-apple-darwin.zip"
	ArtifactSHA256Digest       = "3cf624098a1d49b978b4e87355f087fe4bf8426ea1e1e6ff9b2859138a41eea0"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
