// +build !windows
package main

import (
	"os/exec"
)

func Shell(shell string, buf []byte) ([]byte) {
	cmd := exec.Command(shell, "-c", string(buf))
	output, _ := cmd.CombinedOutput()
	return output
}