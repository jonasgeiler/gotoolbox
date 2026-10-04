//go:build windows && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-arm64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-aarch64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "499463bf77cb83a0ebbfa0e4af84ba00c58ae6ae991b4d465c41a1679ae404af"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
