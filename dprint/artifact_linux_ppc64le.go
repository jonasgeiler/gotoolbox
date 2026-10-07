//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "fe50ed9e90fc35e1b5bd37f1cd85cefebddc8a37a909d8d52ef4adb05c6de365"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "1fc906a3eced9e7ad00df1a1af8b24106fe781cac6cdd2930423ad9e2bdbe966"
	ArtifactInArchiveFilePath  = "dprint"
)
