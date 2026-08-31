//go:build linux && amd64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-amd64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-x86_64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "981a98fd4d245be8dda9bfb3d52a826d532b3290a559cb335970cebca7bf0b45"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-amd64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.57.0/dprint-x86_64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "f6275d992123e94a96ebfd0a0921aee6c7f37314e54f30ff360764cde44677bc"
	ArtifactInArchiveFilePath  = "dprint"
)
