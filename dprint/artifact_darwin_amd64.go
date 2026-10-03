//go:build darwin && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-x86_64-apple-darwin.zip"
	ArtifactSHA256Digest       = "cf3efa1f000125feac934f1afcf76daa35dad4c79dc22ae4c58f92de114e0845"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
