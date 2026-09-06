//go:build darwin && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-aarch64-apple-darwin.zip"
	ArtifactSHA256Digest       = "f9e752812791f258e89ff88a7affd05c793b0db3d929718651eae7df5f19eb2e"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
