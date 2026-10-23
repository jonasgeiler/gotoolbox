//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "b69e7564558b51eb8be1b1f5fb4cc2a7eb87f4e1ac63ec53f58497c4d56d7002"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "d20524979cf322248610f79cec29f4d9c419b84dc49f919960a54edfad03107e"
	ArtifactInArchiveFilePath  = "dprint"
)
