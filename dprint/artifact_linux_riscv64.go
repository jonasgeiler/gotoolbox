//go:build linux && riscv64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = ""
	ArtifactDownloadURL        = ""
	ArtifactSHA256Digest       = ""
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-riscv64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-riscv64gc-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "81b2e4138324b1bb78bafc13e72d8073e001c7164a579ede11128534ec41bfe4"
	ArtifactInArchiveFilePath  = "dprint"
)
