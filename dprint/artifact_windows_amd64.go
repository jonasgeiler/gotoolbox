//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "322e47e188bba6ff3784ef1be0c5f569fc270467eeac4dd10015030599d9466d"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
