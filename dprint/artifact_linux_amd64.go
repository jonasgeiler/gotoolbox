//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "913032d493340643d6b98ce250a2407c5d0abab8613b3e16a09d92661131425a"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "51729ee501593c84e2a2e8233f55959edf2bbd95cbb3998e9f8a81ecad942dba"
	ArtifactInArchiveFilePath  = "dprint"
)
