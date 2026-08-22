package main

import (
	"github.com/jonasgeiler/gotoolbox"
)

var dprint = &gotoolbox.Tool{
	Name:      "dprint",
	Version:   "0.55.2",
	BuildType: gotoolbox.BuildTypeDynamic,
	Binaries: map[gotoolbox.Platform]gotoolbox.DownloadInfo{
		{OS: "android", Arch: "amd64"}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-linux-android.zip",
			Checksum:    "986c5f1926d9a1ae5049ac87c03f0d4cd7fdb823840d47af4fc6080725f3db32",
			ExtractFile: "dprint",
		},
		{OS: "android", Arch: "arm64"}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-linux-android.zip",
			Checksum:    "879a8de2c7e2b17dfd36fff855807803793fefdf8da9b1f69bc07c0a31f94ca3",
			ExtractFile: "dprint",
		},
		{OS: "darwin", Arch: "amd64"}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-apple-darwin.zip",
			Checksum:    "b7074cf6c814f995b783b5baa7e516b34e783c42c9baf7af553dbad731adb3a7",
			ExtractFile: "dprint",
		},
		{OS: "darwin", Arch: "arm64"}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-apple-darwin.zip",
			Checksum:    "e9ba8ed7988f3350501a8cf8af92da616cdec8d9d5c831c069293c587311b49d",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "amd64", Env: gotoolbox.PlatformEnvGNU}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-unknown-linux-gnu.zip",
			Checksum:    "d7ccde62d789dfb048717252d259e21253e32feffe4cbf2dab9954eeab492219",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "amd64", Env: gotoolbox.PlatformEnvMusl}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-unknown-linux-musl.zip",
			Checksum:    "45fc0eef3216af21c4c22c6e7e233aa45c3080fac07b6e94db7008a5c8e5c67a",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "arm64", Env: gotoolbox.PlatformEnvGNU}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-unknown-linux-gnu.zip",
			Checksum:    "299923f2b56d66756ad2c7c220650c72f26437fd3f48b3fb6c0df664073eb1d1",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "arm64", Env: gotoolbox.PlatformEnvMusl}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-unknown-linux-musl.zip",
			Checksum:    "f0101217dd0abc94f1ac01b83d306d0288aeee8a501e8614a5e2bbe037500be0",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "loong64", Env: gotoolbox.PlatformEnvGNU}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-loongarch64-unknown-linux-gnu.zip",
			Checksum:    "9dccf17fd3d79885ece6b3442639cb62cc1ac3852d9a30467fb2f18b4c0997f4",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "loong64", Env: gotoolbox.PlatformEnvMusl}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-loongarch64-unknown-linux-musl.zip",
			Checksum:    "11eb0a855e862bc27c002b863125e96fa794b8d149f919759639fd76d6f31032",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "ppc64le", Env: gotoolbox.PlatformEnvGNU}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-powerpc64le-unknown-linux-gnu.zip",
			Checksum:    "bb7b367b0ad41b413d4b0828cf96d5344e0686e14ab3bb1705a0705b340dd3ec",
			ExtractFile: "dprint",
		},
		{OS: "linux", Arch: "ppc64le", Env: gotoolbox.PlatformEnvMusl}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-powerpc64le-unknown-linux-musl.zip",
			Checksum:    "7973e0203bd5ca23fa50b75c45c61284080500fa50ec45b7ba422bd0f1dad0ed",
			ExtractFile: "dprint",
		},
		// No musl build is available for riscv64.
		{OS: "linux", Arch: "riscv64", Env: gotoolbox.PlatformEnvGNU}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-riscv64gc-unknown-linux-gnu.zip",
			Checksum:    "ed70faf3ecfbb67786470c62fd3eee44172451fa91166f660c7a52a9d9c36979",
			ExtractFile: "dprint",
		},
		// Only MSVC builds are available for Windows.
		{OS: "windows", Arch: "amd64", Env: gotoolbox.PlatformEnvMSVC}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-x86_64-pc-windows-msvc.zip",
			Checksum:    "12e8c26abc8c436223e70f5a30a2864001c92fa356a859eb93e06b97ab7dbd12",
			ExtractFile: "dprint.exe",
		},
		{OS: "windows", Arch: "arm64", Env: gotoolbox.PlatformEnvMSVC}: {
			URL:         "https://github.com/dprint/dprint/releases/download/0.55.2/dprint-aarch64-pc-windows-msvc.zip",
			Checksum:    "65846975b2a8f4e36982ddff875147157c2c9b04c6eb17134d6655ed51e3a931",
			ExtractFile: "dprint.exe",
		},
	},
}

func main() {
	dprint.DownloadAndExec()
}
