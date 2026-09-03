//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "28aa22ea2b009a3d49f9c570bc2a1697af8dff9117cd1cbb7c39deab9f5d5269"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
