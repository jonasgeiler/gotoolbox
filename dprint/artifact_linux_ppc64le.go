//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "d937d7107862d2cd8c3e0c4a094fa3e0fcebe4b534c1a590d4af016edeba463b"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "c99f1599d78199ae51360cd79b16b80d640e81ddf9617de4f95bcbec744b6838"
	ArtifactInArchiveFilePath  = "dprint"
)
