//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "aad2ee102014c563112dcdad65e131dbcfe123355ef948ae91d806e833cdbf40"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "9c2e00cd256dfa2c03a7f568ff08d8277abc5788c7f39fcd9ef325c4642cb4b4"
	ArtifactInArchiveFilePath  = "dprint"
)
