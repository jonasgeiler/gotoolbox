//go:build android && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-android-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-linux-android.zip"
	ArtifactSHA256Digest       = "879a8de2c7e2b17dfd36fff855807803793fefdf8da9b1f69bc07c0a31f94ca3"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
