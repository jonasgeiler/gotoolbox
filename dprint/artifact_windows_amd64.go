//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "2d0be3f0aa8d3c55e4ae19de324353d2179dade2a05b0eda5bd80bc1b94579cc"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
