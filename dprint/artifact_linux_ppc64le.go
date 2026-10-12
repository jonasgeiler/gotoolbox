//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "e8489402c3b2d381e3689d95a35f44b2ebe278827ee8f6bf51ea0ea6f0d60445"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "7d554abef4e360e53eb15ea2815ef3d3158359bfc1e6975247fd54584bc57af9"
	ArtifactInArchiveFilePath  = "dprint"
)
