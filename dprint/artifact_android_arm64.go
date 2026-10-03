//go:build android && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-android-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-aarch64-linux-android.zip"
	ArtifactSHA256Digest       = "3e7d976c0e8cfd088b2ccda2f95e07046341ecf67660ebc5c30835e4c51b4158"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
