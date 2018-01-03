package main

import (
	"log"

	"golang.org/x/crypto/ssh"
)

const (
	USERNAME = "admin"
	PASSWORD = "password123"
	SERVER   = "10.3.2.4:22"
)

func main() {
	c, _ := ssh.Dial(
		"tcp",
		SERVER,
		&ssh.ClientConfig{
			User: USERNAME,
			Auth: []ssh.AuthMethod{
				ssh.Password(PASSWORD),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey,
		},
	)
	if nil != c {
		log.Printf("Up")
	} else {
		log.Printf("Down")
	}
}
