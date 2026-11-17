//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "4aa4207d7739d17870959279652e9c16c2bdadb089bc073343065a7d492afd76"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "549ad89d5047f5dba1f2697f6e64523055ca970faa1251ad794d185140892f70"
	ArtifactInArchiveFilePath  = "dprint"
)
