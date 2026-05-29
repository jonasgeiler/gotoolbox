//go:build !linux

package gotoolbox

// IsMuslLibc returns true if the system is using musl libc.
func IsMuslLibc() bool {
	return false
}
