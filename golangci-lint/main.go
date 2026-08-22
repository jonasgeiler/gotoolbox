package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

var golangciLint = &gotoolbox.Tool{
	Name:    "golangci-lint",
	Version: "v2.12.2",
	Binaries: map[gotoolbox.Platform]gotoolbox.DownloadInfo{
		{OS: "darwin", Arch: "amd64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-darwin-amd64.tar.gz",
			Checksum:    "f6f06d94b6241521c53d15450c5209b028270bf966f842afb11c030c79f5bc16",
			ExtractFile: "golangci-lint-2.12.2-darwin-amd64/golangci-lint",
		},
		{OS: "darwin", Arch: "arm64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-darwin-arm64.tar.gz",
			Checksum:    "a9c54498731b3128f79e090be6110f3e5fffccc617b08142ed244d4126c73f29",
			ExtractFile: "golangci-lint-2.12.2-darwin-arm64/golangci-lint",
		},
		{OS: "linux", Arch: "386"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-386.tar.gz",
			Checksum:    "8acadd219d421b89186438c095fd6da72bcb2cc6a334798d31732003c376233a",
			ExtractFile: "golangci-lint-2.12.2-linux-386/golangci-lint",
		},
		{OS: "linux", Arch: "amd64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-amd64.tar.gz",
			Checksum:    "8df580d2670fed8fa984aac0507099af8df275e665215f5c7a2ae3943893a553",
			ExtractFile: "golangci-lint-2.12.2-linux-amd64/golangci-lint",
		},
		{OS: "linux", Arch: "arm64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-arm64.tar.gz",
			Checksum:    "44cd40a8c76c86755375adfeea52cfd3533cb43d7bd647771e0ae065e166df3a",
			ExtractFile: "golangci-lint-2.12.2-linux-arm64/golangci-lint",
		},
		{OS: "linux", Arch: "armv6"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-armv6.tar.gz",
			Checksum:    "871f97d1a6a8dd8eb2153ec8e1addfc0d2633f42dac1cc8461508a23f971e99d",
			ExtractFile: "golangci-lint-2.12.2-linux-armv6/golangci-lint",
		},
		{OS: "linux", Arch: "armv7"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-armv7.tar.gz",
			Checksum:    "40602c69b04f5262aac21ce090aafb560c4299eadd31dbdc158c074cc4cf9789",
			ExtractFile: "golangci-lint-2.12.2-linux-armv7/golangci-lint",
		},
		{OS: "linux", Arch: "loong64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-loong64.tar.gz",
			Checksum:    "76bfc32dff3597190d1409621c18baa31698c87c52c5b8a7c3c86fdb540c4d73",
			ExtractFile: "golangci-lint-2.12.2-linux-loong64/golangci-lint",
		},
		{OS: "linux", Arch: "mips64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-mips64.tar.gz",
			Checksum:    "d73c73e3f3090659e6ec1276e1f8497d9084690476d7d004672dae1199550b1c",
			ExtractFile: "golangci-lint-2.12.2-linux-mips64/golangci-lint",
		},
		{OS: "linux", Arch: "mips64le"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-mips64le.tar.gz",
			Checksum:    "55429607fb7608f3b1748ece9ab4a74a3eec46ffcaca114bfaf6c0f3d70d4e0d",
			ExtractFile: "golangci-lint-2.12.2-linux-mips64le/golangci-lint",
		},
		{OS: "linux", Arch: "ppc64le"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-ppc64le.tar.gz",
			Checksum:    "31561f2e35ca8e2b9f8c2bc3055c74dd3f0fd341db7c9d0feb5292c95bda1a98",
			ExtractFile: "golangci-lint-2.12.2-linux-ppc64le/golangci-lint",
		},
		{OS: "linux", Arch: "riscv64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-riscv64.tar.gz",
			Checksum:    "32d67a82e5711519aa44ec415e0cb6d1fad9e8d390a95c81e9aeeb1e8a1bf211",
			ExtractFile: "golangci-lint-2.12.2-linux-riscv64/golangci-lint",
		},
		{OS: "linux", Arch: "s390x"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-linux-s390x.tar.gz",
			Checksum:    "e4a35d5531c8c3967e6651f763e371540d2d736972161c2cea35c7601ac0168c",
			ExtractFile: "golangci-lint-2.12.2-linux-s390x/golangci-lint",
		},
		{OS: "windows", Arch: "386"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-windows-386.zip",
			Checksum:    "6242506521a9fba4ba3d86f7d2842d284dcd144ca0f95671ce52c6b0b22a6417",
			ExtractFile: "golangci-lint-2.12.2-windows-386/golangci-lint.exe",
		},
		{OS: "windows", Arch: "amd64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-windows-amd64.zip",
			Checksum:    "bd42e3ebc8cb4ececb86941983baaf1dc221bbb04d838e94ce63b49cc91e02bb",
			ExtractFile: "golangci-lint-2.12.2-windows-amd64/golangci-lint.exe",
		},
		{OS: "windows", Arch: "arm64"}: {
			URL:         "https://github.com/golangci/golangci-lint/releases/download/v2.12.2/golangci-lint-2.12.2-windows-arm64.zip",
			Checksum:    "947b9a5bf762d465710b376c156f0184abb2168378b0826af1899e0ee7183742",
			ExtractFile: "golangci-lint-2.12.2-windows-arm64/golangci-lint.exe",
		},
	},
}

func main() {
	golangciLint.DownloadAndExec()
}
