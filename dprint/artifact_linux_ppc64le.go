//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "718fdb164317347fcec31141926992f083167eb9f54a8df5d4dc77b5c5b46478"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "06a99164542e750085eacbf8a1f984a93defb2d1755ff1a7f12b0d4ab4c1b75c"
	ArtifactInArchiveFilePath  = "dprint"
)
