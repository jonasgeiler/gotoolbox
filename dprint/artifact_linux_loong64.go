//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "848ff9e36d3a16d6f6f660cb8b4df543214d6fd5bd4a6ff1791d3e4719405cf3"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "04a3a540ce665acc3ed7b7304f10259f3c14007211cab9479cb88abbaa29dd39"
	ArtifactInArchiveFilePath  = "dprint"
)
