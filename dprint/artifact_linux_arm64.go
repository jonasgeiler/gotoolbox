//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "24b6f3cf8cc7a116681907843b777a1dc56cd75ffa98727e66a94b2d2f05ad89"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.2/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "ae669413d1996ed8d18e1bc1dfd2b7557064c20820f5c7bcfb9f5e54e47a9b57"
	ArtifactInArchiveFilePath  = "dprint"
)
