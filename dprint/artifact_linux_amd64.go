//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "f105dbc2b0bc0d83c644fcabc3e7bd8779469c0d9faeda48365c58b389d5a443"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.64.3/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "f71cb7e22d6776439eba168d5cb296092d6413fea7148fe4d70684a2702a4fee"
	ArtifactInArchiveFilePath  = "dprint"
)
