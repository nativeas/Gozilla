package main

import (
	"../socket"
	"log"
	"net"
	"os"
)

//实现了一个模拟的客户端
//会往服务端发送4次消息
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	r := socket.NewRemoteObject(conn, 0)
	obj := socket.NewSocketCommand(1, 2, "fff")
	r.Input <- obj
	obj = socket.NewSocketCommand(3, 2, "fff")
	r.Input <- obj
	obj = socket.NewSocketCommand(4, 2, "fff")
	r.Input <- obj
	obj = socket.NewSocketCommand(5, 2, "fff")
	r.Input <- obj
	for {
		_, err := r.Read()
		if err != nil {
			log.Fatal("exit")
		}
	}
}
