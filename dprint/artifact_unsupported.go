//go:build !(android && amd64) && !(android && arm64) && !(darwin && amd64) && !(darwin && arm64) && !(linux && amd64) && !(linux && arm64) && !(linux && loong64) && !(linux && ppc64le) && !(linux && riscv64) && !(windows && amd64) && !(windows && arm64)

package main

//goland:noinspection GoSnakeCaseUsage
const (
	ArtifactCacheName          = ""
	ArtifactDownloadURL        = ""
	ArtifactSHA256Digest       = ""
	ArtifactCacheName_glibc    = ""
	ArtifactDownloadURL_glibc  = ""
	ArtifactSHA256Digest_glibc = ""
	ArtifactInArchiveFilePath  = ""
)
