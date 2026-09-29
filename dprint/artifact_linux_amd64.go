//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "06c2a239a1214d5f9e76c364cadbd0ffaee710fdb516ca6448b08860364db6e2"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "3047dcabe684fe32868d1959f7fad4ccbc2996a393d639bbb6634668fa08dc49"
	ArtifactInArchiveFilePath  = "dprint"
)
