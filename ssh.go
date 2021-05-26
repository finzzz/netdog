package main

import (
	"fmt"
	"log"
	"net"
	"errors"
	"crypto/sha256"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var hash string

func SSHServer(config Config) {
	hash = config.Hash

	privateKey, err := generatePrivateKey()
	if err != nil {
		log.Fatalln(err)
	}

	privateKeyBytes := encodePrivateKeyToPEM(privateKey)
	hostPrivateKeySigner, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalln(err)
	}

	sshConfig := ssh.ServerConfig{
		PasswordCallback: passwordCallback,
	}

	sshConfig.AddHostKey(hostPrivateKeySigner)

	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		SSHServer(config)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		conn.Close()
		SSHServer(config)
		return
	}
	defer conn.Close()

	sshConn, chans, _, err := ssh.NewServerConn(conn, &sshConfig)
	if err != nil {
		SSHServer(config)
		return
	}
	defer sshConn.Close()

	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		channel, _, err := newChannel.Accept()
		if err != nil {
			continue
		}
		
		term := terminal.NewTerminal(channel, "> ")

		go func() {
			defer channel.Close()
			for {
				cmd, err := term.ReadLine()
				if err != nil {
					break
				}

				output := Shell(config.Shell, []byte(cmd))
				term.Write(output)
			}
		}()
	}

	SSHServer(config)
}

func generatePrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

func passwordCallback(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	if fmt.Sprintf("%x",sha256.Sum256(password)) == hash {
		return nil, nil
	}

	return nil, errors.New("Auth Failed")
}