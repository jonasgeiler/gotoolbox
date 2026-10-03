//go:build linux && riscv64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = ""
	ArtifactDownloadURL        = ""
	ArtifactSHA256Digest       = ""
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-riscv64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-riscv64gc-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "b88f8c1d6affa3743f60404feed51703b8d77306d0fe3b451e4fb7a4e4f84054"
	ArtifactInArchiveFilePath  = "dprint"
)
