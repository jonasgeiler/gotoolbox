//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "359ed8678d9660b8e5016fe4bef111f23b7fcff890abd3deab02f8720b61d482"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "016a52d84f20c82b074d9dfbd0d802193fc2f3b97a86ff4c06a50f6a1a1c91c1"
	ArtifactInArchiveFilePath  = "dprint"
)
