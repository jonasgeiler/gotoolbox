//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "5ea6e8459078d9bf1959fe2362cabdcb317e2cd7ef44e70e29e3300e623f6fba"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "b93091e0113d877dd4cc72ae161d63bea7e02407008c3bb1e2442f9dd11f6a22"
	ArtifactInArchiveFilePath  = "dprint"
)
