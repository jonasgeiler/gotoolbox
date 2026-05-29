//go:build !unix

package gotoolbox

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec(binPath string) {
	cmd := exec.Command(binPath, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		if exit, ok := err.(*exec.ExitError); ok {
			os.Exit(exit.ExitCode())
			return
		}
		fmt.Fprintf(os.Stderr, "Exec Error: executing %s: %v\n", binPath, err)
		os.Exit(1)
		return
	}
	os.Exit(0)
}
