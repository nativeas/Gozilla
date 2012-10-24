package main

import (
	"../socket"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	r := socket.NewRemoteObject(conn)
	obj := socket.NewSocketCommand(1, 2, "fff")
	r.Input <- obj
	obj = socket.NewSocketCommand(3, 2, "fff")
	r.Input <- obj
	obj = socket.NewSocketCommand(4, 2, "fff")
	r.Input <- obj
	obj = socket.NewSocketCommand(5, 2, "fff")
	r.Input <- obj
	for {
		r.Read()
	}
}
