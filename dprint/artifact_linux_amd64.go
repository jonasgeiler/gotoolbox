//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "03f3e8002f9d952bf53325fc8853686d2e8808b9e5b652b215600ba147e73a5c"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.4/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "1d26357d8bc66898d4ecc7dafca3d2971cfe222ba487228bd2ce5c3b3a309c33"
	ArtifactInArchiveFilePath  = "dprint"
)
