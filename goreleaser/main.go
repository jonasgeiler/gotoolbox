package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

var goreleaser = &gotoolbox.Tool{
	Name:    "goreleaser",
	Version: "v2.16.0",
	Binaries: map[gotoolbox.Platform]gotoolbox.DownloadInfo{
		{OS: "darwin", Arch: "amd64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Darwin_x86_64.tar.gz",
			Checksum:    "2b82d8319ee517d4242b48a858114b267c621f1dd1fe51a14680902b18a5dac8",
			ExtractFile: "goreleaser",
		},
		{OS: "darwin", Arch: "arm64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Darwin_arm64.tar.gz",
			Checksum:    "8f6898256f35531165d90f2db581c5ee0d32bda83ebc25ac231ff5bdb9d2071a",
			ExtractFile: "goreleaser",
		},
		{OS: "linux", Arch: "386"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_i386.tar.gz",
			Checksum:    "b6b0764b9e339fcfd8bcf1786424c99ddecbecf27d15025c189e1c64932a1563",
			ExtractFile: "goreleaser",
		},
		{OS: "linux", Arch: "amd64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_x86_64.tar.gz",
			Checksum:    "eaae05b5eba07533bd0f06846b68c808399504784df00c62eb219541fc04e5e2",
			ExtractFile: "goreleaser",
		},
		{OS: "linux", Arch: "arm"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_armv7.tar.gz",
			Checksum:    "c65b905052a85a2f3248bc85ce77cc62ac600302a595b524d375b850f85e8958",
			ExtractFile: "goreleaser",
		},
		{OS: "linux", Arch: "arm64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_arm64.tar.gz",
			Checksum:    "0102d974373fcdeb77042d1f5897caffa193be36620fdc6c1da43a01ef8e10d3",
			ExtractFile: "goreleaser",
		},
		{OS: "linux", Arch: "loong64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_loong64.tar.gz",
			Checksum:    "826f70d2f225e44b295a710ae229aa79f1ee5ef10d61cd537c1cf07113196060",
			ExtractFile: "goreleaser",
		},
		{OS: "linux", Arch: "ppc64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_ppc64.tar.gz",
			Checksum:    "34920ed822616e10216069fec380c832ffad0501d7a9de6680aa103169e940b6",
			ExtractFile: "goreleaser",
		},
		{OS: "linux", Arch: "riscv64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Linux_riscv64.tar.gz",
			Checksum:    "d3bfa4f1f4639d45cac045dd129dfabce6ece2b198de0391531b464f67cae273",
			ExtractFile: "goreleaser",
		},
		{OS: "windows", Arch: "386"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Windows_i386.zip",
			Checksum:    "d5b702f899a357670e1bd6ecc692a35baa61974185adb968c51ca441776cab7f",
			ExtractFile: "goreleaser.exe",
		},
		{OS: "windows", Arch: "amd64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Windows_x86_64.zip",
			Checksum:    "6fe5eda11f0bcac8069aff3ef3dcb0b11816c9e95f89773595564697a5278bc9",
			ExtractFile: "goreleaser.exe",
		},
		{OS: "windows", Arch: "arm64"}: {
			URL:         "https://github.com/goreleaser/goreleaser/releases/download/v2.16.0/goreleaser_Windows_arm64.zip",
			Checksum:    "1183c81863044ce9baa89c1393c258949390b8df683df7ca959e9c718d7723c9",
			ExtractFile: "goreleaser.exe",
		},
	},
}

func main() {
	goreleaser.DownloadAndExec()
}
