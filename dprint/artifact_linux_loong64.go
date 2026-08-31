//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "39e2bbe47657b411c2f7b0a2ced3ff120fdcbbf5023cca9904e0be262f558a16"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "7d215fd58795d4adabcfab32344a8eca8a25d056514c53695ae6f5eb4bf9bf2b"
	ArtifactInArchiveFilePath  = "dprint"
)
