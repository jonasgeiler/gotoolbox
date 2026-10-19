//go:build android && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-android-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-x86_64-linux-android.zip"
	ArtifactSHA256Digest       = "7dc159e4752207195574f8c32f08ad1e02f44a2bdf3ca5fd0c78393910ecebc2"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
