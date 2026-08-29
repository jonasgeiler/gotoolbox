//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "12e8c26abc8c436223e70f5a30a2864001c92fa356a859eb93e06b97ab7dbd12"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
