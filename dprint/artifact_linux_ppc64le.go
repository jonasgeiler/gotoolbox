//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "3bcaec88382d72d40c774f8e610371d3842b72497d179723fcbe1dfed09040e1"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.65.0/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "66de75206b1b5605a9c9f22d5e74a6ae140d78e6886977d1402d3243c43160b6"
	ArtifactInArchiveFilePath  = "dprint"
)
