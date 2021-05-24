// +build windows
package main

import (
	"syscall"
	"os/exec"
)

func Shell(shell string, buf []byte) ([]byte) {
	cmd := exec.Command(shell, "-c", string(buf))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, _ := cmd.CombinedOutput()
	return output
}