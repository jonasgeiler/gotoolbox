package main

import (
	"crypto"

	"github.com/jonasgeiler/gotoolbox"
)

var Tool = &gotoolbox.Tool{
	Name: "golangci-lint",
	Artifact: gotoolbox.Artifact{
		CacheName:   ArtifactCacheName,
		DownloadURL: ArtifactDownloadURL,
		Checksum: gotoolbox.Checksum{
			Algorithm: crypto.SHA256,
			Digest:    ArtifactSHA256Digest,
		},
		ArchiveFormat:     ArtifactArchiveFormat,
		InArchiveFilePath: ArtifactInArchiveFilePath,
	},
}

func main() {
	Tool.Run()
}
