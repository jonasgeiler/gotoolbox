//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "8914a940bdf4b9ba438de9ccfb4317105a8501eaf9d5593784c9f41387243475"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "f58b5d7c13ccd613cfa264f1ae82d4e5db1373f8fe20462bc1126051d111b0ce"
	ArtifactInArchiveFilePath  = "dprint"
)
