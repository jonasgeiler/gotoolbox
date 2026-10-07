//go:build linux && riscv64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = ""
	ArtifactDownloadURL        = ""
	ArtifactSHA256Digest       = ""
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-riscv64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-riscv64gc-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "c5af6a12fbecd25c858af69910dc0f4791552c913b8b73293d44ab00d6a1b71b"
	ArtifactInArchiveFilePath  = "dprint"
)
