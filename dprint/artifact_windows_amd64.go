//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "6aa57ba5f21a6247f2a3144c2e16ee7dd9aa1f805850d65f8a554cb38b575584"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
