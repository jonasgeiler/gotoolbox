//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "21ac931303bb420bbdf3e0589506e5c32611c9fcff4ed2eb9802db657c499e9a"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "5102aa940a060e48c33a5ceefce1273c701c6618503b7514323caada3b3649bb"
	ArtifactInArchiveFilePath  = "dprint"
)
