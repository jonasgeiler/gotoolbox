//go:build windows && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-arm64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-aarch64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "913ceec8446d297d33f74c37dbddec0e5292b9a9da99bc7328bdd0c3144c7b9e"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
