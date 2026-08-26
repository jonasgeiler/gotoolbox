//go:build linux

package gotoolbox

import (
	"os"
	"os/exec"
	"regexp"
	"sync"
)

var muslRegex = regexp.MustCompile(`\bmusl\b`)
var glibcRegex = regexp.MustCompile(`\bGNU +libc\b`)

// IsGlibcEnv returns true if the host platform is a GNU libc environment.
// Otherwise, e.g. on musl libc environments, returns false.
var IsGlibcEnv = sync.OnceValue(
	func() bool {
		// TODO: Detect with CGO and preprocessor definitions?

		if out, err := os.ReadFile("/proc/self/maps"); err == nil {
			if muslRegex.Match(out) {
				return false
			}
		}

		if lddPath, err := exec.LookPath("ldd"); err == nil {
			if ldd, err := os.ReadFile(lddPath); err == nil {
				if muslRegex.Match(ldd) {
					return false
				} else if glibcRegex.Match(ldd) {
					return true
				}
			}

			if lddVersion, err := exec.Command(
				lddPath, "--version",
			).CombinedOutput(); err == nil {
				if muslRegex.Match(lddVersion) {
					return false
				} else if glibcRegex.Match(lddVersion) {
					return true
				}
			}
		}

		return false
	},
)
