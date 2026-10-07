//go:build android && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-android-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-x86_64-linux-android.zip"
	ArtifactSHA256Digest       = "7340aae74956c5e8da20904d8582e0ca4f51df305acd745add0fd7041e41d0de"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
