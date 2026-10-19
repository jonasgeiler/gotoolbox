//go:build windows && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-arm64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-aarch64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "55a758b94fbe009acae7aefa7638c4820709e19b0bfaf268405c11af54f56f2b"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
