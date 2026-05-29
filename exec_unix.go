//go:build unix

package gotoolbox

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func Exec(binPath string) {
	// On Unix, we can use syscall.Exec to replace the current process instead
	// of spawning a sub-process, which is more efficient.
	argv := os.Args
	argv[0] = filepath.Base(binPath)
	if err := syscall.Exec(binPath, argv, os.Environ()); err != nil {
		fmt.Fprintf(os.Stderr, "Exec Error: executing %s: %v\n", binPath, err)
		os.Exit(1)
	}
}
