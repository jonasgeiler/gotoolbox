//go:build android && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-android-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-x86_64-linux-android.zip"
	ArtifactSHA256Digest       = "f85b1998119f0d369397828a8c1262160b27d9b9b92ca75ab7952e444c344e8b"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
