//go:build !(darwin && amd64) && !(darwin && arm64) && !(linux && 386) && !(linux && amd64) && !(linux && arm64) && !(linux && arm.6) && !(linux && arm.7) && !(linux && loong64) && !(linux && mips64) && !(linux && mips64le) && !(linux && ppc64le) && !(linux && riscv64) && !(linux && s390x) && !(windows && 386) && !(windows && amd64) && !(windows && arm64)

package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

const (
	ArtifactCacheName         = ""
	ArtifactDownloadURL       = ""
	ArtifactSHA256Digest      = ""
	ArtifactArchiveFormat     = gotoolbox.NotAnArchive
	ArtifactInArchiveFilePath = ""
)
