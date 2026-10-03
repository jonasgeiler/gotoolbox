//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "b100406bc5d3098cd3756b1e7cc7e7c78e05a13e88c4d1593ebb1d3e2a95b26a"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.59.0/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "fbf544f4149e5a6ed7d5feefbb85ee7570cf58fc4dc79e7e6245293669f4f34f"
	ArtifactInArchiveFilePath  = "dprint"
)
