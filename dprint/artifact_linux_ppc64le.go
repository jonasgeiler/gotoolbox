//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "e57e201bc9f28937f65f2a94a5c08ec72386f93a8b916429c08cf2e4e848b493"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.56.1/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "a264d977e5b0ead41d9149855aad710786eb1ff3737351f118c9d834f4e74982"
	ArtifactInArchiveFilePath  = "dprint"
)
