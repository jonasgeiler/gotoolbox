//go:build windows && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-arm64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-aarch64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "de86d320f0753f8803ef957baec766be18858e4021ddcf2d0584142a8ace1bae"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
