//go:build !linux

package gotoolbox

// IsHostPlatformEnvMusl returns true if the host platform is a musl libc environment.
func IsHostPlatformEnvMusl() bool {
	return false
}
