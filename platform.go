package gotoolbox

import (
	"runtime"
	"strings"
)

// PlatformEnv identifies the C library or ABI environment that a binary
// is compatible with, such as GNU libc, musl libc, or the MSVC ABI.
type PlatformEnv int

const (
	// PlatformEnvAny indicates that the binary is not specific to a C library
	// or ABI environment.
	PlatformEnvAny PlatformEnv = iota

	// PlatformEnvGNU indicates a GNU libc (glibc) environment.
	PlatformEnvGNU

	// PlatformEnvMusl indicates a musl libc environment.
	PlatformEnvMusl

	// PlatformEnvMSVC indicates an MSVC ABI environment on Windows.
	PlatformEnvMSVC
)

// String returns a human-readable string representation of PlatformEnv.
func (e PlatformEnv) String() string {
	switch e {
	case PlatformEnvGNU:
		return "glibc"
	case PlatformEnvMusl:
		return "musl libc"
	case PlatformEnvMSVC:
		return "MSVC"
	default:
		return ""
	}
}

// DirName returns a string representation of PlatformEnv that is usable as the
// name of a directory. It considers an empty string a valid directory name,
// since filepath.Join ignores empty strings when building a path.
func (e PlatformEnv) DirName() string {
	switch e {
	case PlatformEnvGNU:
		return "gnu"
	case PlatformEnvMusl:
		return "musl"
	case PlatformEnvMSVC:
		return "msvc"
	default:
		return ""
	}
}

// Platform identifies an operating system, CPU architecture, and C library or
// ABI environment for which a binary is built or on which it can run.
type Platform struct {
	OS, Arch string
	Env      PlatformEnv
}

// String returns a human-readable string representation of Platform.
func (p Platform) String() string {
	b := strings.Builder{}
	b.WriteString(p.OS)
	b.WriteString("/")
	b.WriteString(p.Arch)
	if p.Env != PlatformEnvAny {
		b.WriteString(" (")
		b.WriteString(p.Env.String())
		b.WriteString(")")
	}
	return b.String()
}

// hostPlatformEnv detects and returns the C library or ABI environment of the
// host platform.
func hostPlatformEnv() PlatformEnv {
	switch runtime.GOOS {
	case "linux":
		if IsHostPlatformEnvMusl() {
			return PlatformEnvMusl
		}
		return PlatformEnvGNU

	case "windows":
		// Always uses MSVC at the moment.
		return PlatformEnvMSVC

	default:
		return PlatformEnvAny
	}
}

// HostPlatform returns the platform of the host machine. Used to determine the
// correct binary when downloading a tool.
//
// When withEnv is set to true, it will also try to detect the environment of
// the host platform.
func HostPlatform(withEnv bool) Platform {
	if withEnv {
		return Platform{
			OS:   runtime.GOOS,
			Arch: runtime.GOARCH,
			Env:  hostPlatformEnv(),
		}
	}

	return Platform{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
}
