//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "e72c698a33924b42eb9926b99d30378a95088ee4f4dcb026c867d53590f3c4c5"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "2da200b8c59763769a1b996de1f71821e79eb70be680b2c70c313febad26151e"
	ArtifactInArchiveFilePath  = "dprint"
)
