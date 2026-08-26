//go:build linux

package gotoolbox

import (
	"os"
	"os/exec"
	"regexp"
	"sync"
)

var muslRegex = regexp.MustCompile(`\bmusl\b`)
var gnuLibcRegex = regexp.MustCompile(`\bGNU +libc\b`)

// IsHostPlatformEnvMusl returns true if the host platform is a musl libc environment.
var IsHostPlatformEnvMusl = sync.OnceValue(
	func() bool {
		if out, err := os.ReadFile("/proc/self/maps"); err == nil {
			if muslRegex.Match(out) {
				return true
			}
		}

		if lddPath, err := exec.LookPath("ldd"); err == nil {
			if ldd, err := os.ReadFile(lddPath); err == nil {
				if muslRegex.Match(ldd) {
					return true
				} else if gnuLibcRegex.Match(ldd) {
					return false
				}
			}

			if lddVersion, err := exec.Command(
				lddPath, "--version",
			).CombinedOutput(); err == nil {
				if muslRegex.Match(lddVersion) {
					return true
				} else if gnuLibcRegex.Match(lddVersion) {
					return false
				}
			}
		}

		return false
	},
)
