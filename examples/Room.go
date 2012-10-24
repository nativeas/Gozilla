package main

import (
	"../socket"
	"log"
)

func main() {
	s := socket.NewRemoteRoom("127.0.0.1:8080")
	log.Print(s)
	s.StartDaemon()
}
