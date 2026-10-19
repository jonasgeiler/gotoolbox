//go:build android && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-android-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-aarch64-linux-android.zip"
	ArtifactSHA256Digest       = "17ddb34b9e452bb1ef79168ffb85c5780b58612ffc666c5620257ca3ce6d4b8f"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
