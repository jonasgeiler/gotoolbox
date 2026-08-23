//go:build !(darwin && amd64) && !(darwin && arm64) && !(linux && 386) && !(linux && amd64) && !(linux && arm) && !(linux && arm64) && !(linux && loong64) && !(linux && ppc64) && !(linux && riscv64) && !(windows && 386) && !(windows && amd64) && !(windows && arm64)

package main

const (
	DownloadURL = ""
	SHA256Sum   = ""
	ExtractFile = ""
)
