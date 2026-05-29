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

// IsMuslLibc returns true if the system is using musl libc.
var IsMuslLibc = sync.OnceValue(
	func() bool {
		out, err := os.ReadFile("/proc/self/maps")
		if err == nil {
			if muslRegex.Match(out) {
				return true
			}
		}

		lddPath, err := exec.LookPath("ldd")
		if err == nil {
			ldd, err := os.ReadFile(lddPath)
			if err == nil {
				if muslRegex.Match(ldd) {
					return true
				} else if gnuLibcRegex.Match(ldd) {
					return false
				}
			}

			lddVersion, err := exec.
				Command(lddPath, "--version").
				CombinedOutput()
			if err == nil {
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
