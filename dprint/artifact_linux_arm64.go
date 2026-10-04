//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "9b537fdfb75b2402f47d15cc615f1ba31736a37e2aed65662e8d9d4b647c70f1"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "03fd8908f9b418bdc4701347ff0527a59d1c08db2bc09a9f5d8101eb711e7fe2"
	ArtifactInArchiveFilePath  = "dprint"
)
