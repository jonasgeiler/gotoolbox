//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "57562ea5e06d07a04a85ecb90e5827db54e551f8bd0ebc52b1fb44f7e851c027"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "79ff42f3978b6606bfb64a7d869a9d1923f0b95abec1419848bb724e484295b0"
	ArtifactInArchiveFilePath  = "dprint"
)
