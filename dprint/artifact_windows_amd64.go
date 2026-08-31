//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "d2c5ee170ff88c0b16ea1f5f214e4d74877c7edb9ca701da5c9851c0c1ef9338"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
