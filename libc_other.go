//go:build !linux

package gotoolbox

// IsGlibcEnv returns true if the host platform is a GNU libc environment.
// Otherwise, e.g. on musl libc environments, returns false.
func IsGlibcEnv() bool {
	return false
}
