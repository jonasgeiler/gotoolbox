//go:build darwin && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-darwin-arm64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-aarch64-apple-darwin.zip"
	ArtifactSHA256Digest       = "3fbf19db032dea2e8eb6e52eebbc8c328954a46d52ffd97d99a95b9c9b81414f"
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
