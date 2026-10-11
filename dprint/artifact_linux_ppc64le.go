//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "eb43af95467f0f955392c5cc390deaf2f85e5b172150f3ca13ed91ff8769a940"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "bbb1e4430d0f9bf7613dfdd979c415f835053b6240453112f97260e0f4592b08"
	ArtifactInArchiveFilePath  = "dprint"
)
