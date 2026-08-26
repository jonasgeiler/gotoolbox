//go:build unix

package gotoolbox

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func Exec(binPath string) error {
	// On Unix, we can use syscall.Exec to replace the current process instead
	// of spawning a sub-process, which is more efficient.
	argv := os.Args
	argv[0] = filepath.Base(binPath)
	if err := syscall.Exec(binPath, argv, os.Environ()); err != nil {
		return fmt.Errorf("executing %s: %w", binPath, err)
	}
	return nil
}
