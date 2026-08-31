//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "239999452ff79f6325e178be69913dd95ed428c05d1e8840e53ae42d5eed268d"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "b754b6fdc4f4f3ebdf53b02b41eb75ca272a7f641ac508913d595446453fdb56"
	ArtifactInArchiveFilePath  = "dprint"
)
