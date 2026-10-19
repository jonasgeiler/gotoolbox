//go:build darwin && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-x86_64-apple-darwin.zip"
	ArtifactSHA256Digest       = "c5eb2a8b07e12d32818ee9f3c19c708d88be2c974eedaab7c4a31e902e30f429"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
