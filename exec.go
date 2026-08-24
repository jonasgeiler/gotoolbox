//go:build !unix

package gotoolbox

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func Exec(binPath string) error {
	cmd := exec.Command(binPath, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		if exit, ok := errors.AsType[*exec.ExitError](err); ok {
			// The program exited with non-zero exit code, but the command
			// was successful, so just exit with the program's exit code.
			os.Exit(exit.ExitCode())
			return nil
		}
		return fmt.Errorf("executing %s: %w", binPath, err)
	}
	os.Exit(0)
	return nil
}
