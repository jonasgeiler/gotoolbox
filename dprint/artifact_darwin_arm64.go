//go:build darwin && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-aarch64-apple-darwin.zip"
	ArtifactSHA256Digest       = "1888b152b3541c3690ec5291dd58bde7116e6c8beb4aa1f242187f22071df22f"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
