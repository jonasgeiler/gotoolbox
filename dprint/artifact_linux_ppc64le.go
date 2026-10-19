//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "89fb5f4876e5db8fc8d369934e491c4856ebdef047e516f2b4e0f2cee53d3ef4"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.0/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "19f6ae54f02301e30fc6b4c6075a80b86b32e8429b4266ca2c7aa710327bcc62"
	ArtifactInArchiveFilePath  = "dprint"
)
