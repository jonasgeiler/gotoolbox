//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "dd0fc5214329c95e552d7bfbb86d4a608e64f9482c6ec33dd182f4c839ee11cd"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "27f2f0ef079307068ae49a12bbfe09a4c29d278a9316d2b222cc0969a57cea03"
	ArtifactInArchiveFilePath  = "dprint"
)
