//go:build linux && riscv64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = ""
	ArtifactDownloadURL        = ""
	ArtifactSHA256Digest       = ""
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-riscv64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-riscv64gc-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "ff5d6b45c9fb1f49934af33ae20bec0c0655ccfdbce0a9a417ee706bc4171d6d"
	ArtifactInArchiveFilePath  = "dprint"
)
