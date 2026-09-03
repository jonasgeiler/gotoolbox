//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "b9f1d4f89b51f3d291d2be61132f81f451d06e9a3e33aab2510b69564a8d3648"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "a7d21498b38598c9898e61b1dc8a37cdaab15499c58390bbf9100da1864342f7"
	ArtifactInArchiveFilePath  = "dprint"
)
