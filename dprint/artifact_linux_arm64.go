//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v0.55.2-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "f0101217dd0abc94f1ac01b83d306d0288aeee8a501e8614a5e2bbe037500be0"
	ArtifactCacheName_glibc    = "dprint-v0.55.2-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "299923f2b56d66756ad2c7c220650c72f26437fd3f48b3fb6c0df664073eb1d1"
	ArtifactInArchiveFilePath  = "dprint"
)
