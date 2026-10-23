//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "61c8936368ff4030a615ce43853adc5a0707ff77973743d0a691aee2cbf64723"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "c820bfc67b0cd6f5f0fa4d00ecdcb056944c86030859f5e50151c88d9164b30d"
	ArtifactInArchiveFilePath  = "dprint"
)
