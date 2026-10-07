//go:build windows && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-arm64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-aarch64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "061f0c0790de5e6758c73c0d2639f698b9802c1bb5c6c70124c5ebb112b1afce"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
