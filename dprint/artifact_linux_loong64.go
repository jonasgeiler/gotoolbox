//go:build linux && loong64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-loong64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-loongarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "cf7af7c011e29477569523a8c3ea17167a844c3e536c134319ec30d818eb8d73"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-loong64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.61.1/dprint-loongarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "fc015e872f00d89675d90ed950c6e91de9c712580dab5077be6f41b0caf5a96f"
	ArtifactInArchiveFilePath  = "dprint"
)
