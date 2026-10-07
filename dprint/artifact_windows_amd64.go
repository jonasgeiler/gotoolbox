//go:build windows && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-windows-amd64-msvc"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-x86_64-pc-windows-msvc.zip"
	ArtifactSHA256Digest       = "f259df6beed31a7242a613cf1ff948637ca17d1b385f8effa2cfbf6a056ff84c"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint.exe"
)
