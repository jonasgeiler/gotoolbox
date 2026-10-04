//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "022e540ad9c45e2ad8d5e95c87543316d6b19af4b6f51697d0731885c8add339"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.60.1/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "f0ab7c56d08611740c346a53185488ee237ab56ac328283b4e3b456f821b71aa"
	ArtifactInArchiveFilePath  = "dprint"
)
