//go:build darwin && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-x86_64-apple-darwin.zip"
	ArtifactSHA256Digest       = "b94e2ede59904f34f8047ee9f069767acf19df389b3983f4b97a4fd64f5fcf15"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
