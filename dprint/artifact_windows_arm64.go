//go:build windows && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-arm64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-aarch64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "5bce8673fe5ff8c5162d8a51fef54cd98c1213e9d4b37be03bca6cd21b196b09"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
