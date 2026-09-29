//go:build linux && ppc64le

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-ppc64le-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-powerpc64le-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "b9616bd309b562ab2a263b2c8555e65237c8c9db4ad6936cff1c04b1173b6bef"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-ppc64le-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.58.0/dprint-powerpc64le-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "1cf739e5b89d6a82ea16c42131aca07e767ff086437262680310b905f41bb271"
	ArtifactInArchiveFilePath  = "dprint"
)
