//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "a7e1cba0b951bce0355e18f9045fe461244b6c022e5bd2a2061616aa14530880"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "5232a896fb1b957c9cbf589d4581077f1ea8d22cdcb20284d399a37b0d26ca77"
	ArtifactInArchiveFilePath  = "dprint"
)
