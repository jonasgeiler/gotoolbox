//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "d956144fa8d873dc43c37d1811b3f8c42701f74bb0be7c25c543b189ab53d6c0"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.1/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "4117443cc5fade617dd5f4850cc2fa6b2902136052ef313ee20947036c3554cc"
	ArtifactInArchiveFilePath  = "dprint"
)
