//go:build windows && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v0.55.2-windows-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "65846975b2a8f4e36982ddff875147157c2c9b04c6eb17134d6655ed51e3a931"
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
