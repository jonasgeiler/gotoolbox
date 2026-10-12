//go:build linux && arm64

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = "dprint-v" + Version + "-linux-arm64-musl"
	ArtifactDownloadURL        = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-aarch64-unknown-linux-musl.zip"
	ArtifactSHA256Digest       = "26df4a1534db02bb1c1f614e91f02fcee786357de80094a1fd5dcce27dec803e"
	ArtifactCacheName_glibc    = "dprint-v" + Version + "-linux-arm64-glibc"
	ArtifactDownloadURL_glibc  = "https://github.com/dprint/dprint/releases/download/0.63.0/dprint-aarch64-unknown-linux-gnu.zip"
	ArtifactSHA256Digest_glibc = "8dfbf17fc4be97020e51bd47e1370acdb019301295ed539840016671cab33e89"
	ArtifactInArchiveFilePath  = "dprint"
)
