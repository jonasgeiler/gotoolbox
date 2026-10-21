//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "bd6217389e91b580517183b2b7ebdf15ffe4f779b353f1714c5cb6b0d444cf30"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.1/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "f6196308d7566f22a03211e102779ae1d7fcf4ea4e8ab33030262cd23d29afce"
	ArtifactInArchiveFilePath  = "dprint"
)
