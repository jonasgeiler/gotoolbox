//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "1038af32fade7a79f9c3a690d9546bb17be13dd7b4568a3684692fdcb0a52a1d"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
