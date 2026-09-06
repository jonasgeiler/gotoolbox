//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "2ff5a32f8d2f641cb252636c79972e59db0a7a40f11699cbbddd76cbb063a31c"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "e69a55aa4b242a1870b2a1e39ba03239950ea956a4fab886dc6580c41c29644a"
	ArtifactInArchiveFilePath  = "dprint"
)
