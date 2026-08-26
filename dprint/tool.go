package main

import (
	"crypto"
	_ "crypto/sha256"

	"github.com/jonasgeiler/gotoolbox"
)

var Tool = &gotoolbox.Tool{
	Name: "dprint",
	Artifact: gotoolbox.Artifact{
		CacheName: gotoolbox.SelectByLibc(
			ArtifactCacheName_glibc,
			ArtifactCacheName,
		),
		DownloadURL: gotoolbox.SelectByLibc(
			ArtifactDownloadURL_glibc,
			ArtifactDownloadURL,
		),
		Checksum: gotoolbox.Checksum{
			Algorithm: crypto.SHA256,
			Digest: gotoolbox.SelectByLibc(
				ArtifactSHA256Digest_glibc,
				ArtifactSHA256Digest,
			),
		},
		ArchiveFormat:     gotoolbox.ZipArchive,
		InArchiveFilePath: ArtifactInArchiveFilePath,
	},
}

func main() {
	Tool.Run()
}
