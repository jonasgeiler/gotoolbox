//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "5f61c6898dfb3db8e54f9405727d168eb512a403d941ddd5b725c328bb8d5bc6"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "147ca90b1348036344c9e4e45f0985acb37c18f36a9d3a1f37d4f553ebb80d58"
	ArtifactInArchiveFilePath  = "dprint"
)
