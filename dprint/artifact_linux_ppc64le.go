//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "7973e0203bd5ca23fa50b75c45c61284080500fa50ec45b7ba422bd0f1dad0ed"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "bb7b367b0ad41b413d4b0828cf96d5344e0686e14ab3bb1705a0705b340dd3ec"
	ArtifactInArchiveFilePath  = "dprint"
)
