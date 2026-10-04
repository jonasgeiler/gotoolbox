//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "c4e5e8491f52971f8e3bd0ed80995eafd9119fbe603ba626694a20d56c14a730"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "5612c587a34dfbbda37162d46755f9e8c6a4efba7950920b2a308ea84ce0886a"
	ArtifactInArchiveFilePath  = "dprint"
)
