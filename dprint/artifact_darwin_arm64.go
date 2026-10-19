//go:build darwin && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-aarch64-apple-darwin.zip"
	ArtifactSHA256Digest       = "791560cf48299933a9fff2aea93b7e42b39e62d8a109634619d01c4f2c1aa68a"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
