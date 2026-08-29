//go:build linux && riscv64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = ""
	ArtifactDownloadURL        = ""
	ArtifactSHA256Digest       = ""
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-riscv64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-riscv64gc-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "ed70faf3ecfbb67786470c62fd3eee44172451fa91166f660c7a52a9d9c36979"
	ArtifactInArchiveFilePath  = "dprint"
)
