//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "e8d4a78fc6aa1b1a3f007d947a3c9a1ae1b0c240868107c496eb2992cc1630be"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.62.0/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "9d2de0313bbfc26ed80148a38a430c037a216a905cdf317f6271f8a10311933a"
	ArtifactInArchiveFilePath  = "dprint"
)
