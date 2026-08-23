//go:build !(darwin && amd64) && !(darwin && arm64) && !(linux && 386) && !(linux && amd64) && !(linux && arm64) && !(linux && armv6) && !(linux && armv7) && !(linux && loong64) && !(linux && mips64) && !(linux && mips64le) && !(linux && ppc64le) && !(linux && riscv64) && !(linux && s390x) && !(windows && 386) && !(windows && amd64) && !(windows && arm64)

package main

const (
	DownloadURL = ""
	SHA256Sum   = ""
	ExtractFile = ""
)
