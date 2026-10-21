//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "cc5a9b124b9305cdbba14ebd32e45dcf76d213f4ded7ee73a87f5bfae176bcc7"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "8a0c9a1a8558db8d99749533f35a6560196765ce6f0b02003ca7797067668f5b"
	ArtifactInArchiveFilePath  = "dprint"
)
