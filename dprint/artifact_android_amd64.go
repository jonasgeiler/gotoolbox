//go:build android && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-android-amd64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-linux-android.zip"
	ArtifactSHA256Digest       = "986c5f1926d9a1ae5049ac87c03f0d4cd7fdb823840d47af4fc6080725f3db32"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
