//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v0.55.2-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "45fc0eef3216af21c4c22c6e7e233aa45c3080fac07b6e94db7008a5c8e5c67a"
	ArtifactCacheName_glibc    = "dprint-v0.55.2-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "d7ccde62d789dfb048717252d259e21253e32feffe4cbf2dab9954eeab492219"
	ArtifactInArchiveFilePath  = "dprint"
)
