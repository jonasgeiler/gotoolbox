//go:build darwin && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-aarch64-apple-darwin.zip"
	ArtifactSHA256Digest       = "06a7ea017e4d7da296f8a44fabb2f3d6ffafa91bbd5fddd4830bdc418bacb46d"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
