//go:build linux && riscv64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v0.55.2-linux-riscv64"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-riscv64gc-unknown-linux-gnu.zip"
	ArtifactSHA256Digest       = "ed70faf3ecfbb67786470c62fd3eee44172451fa91166f660c7a52a9d9c36979"
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = "dprint"
)
