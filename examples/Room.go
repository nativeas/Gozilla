package main

import (
	"../socket"
	"log"
)

// 实现了一个服务器的机制
// s.Output用来切换到逻辑线程
func main() {
	s := socket.NewRemoteRoom("127.0.0.1:8080")
	log.Print(s)
	go s.StartDaemon()
	for {
		select {
		case obj := <-s.Output:
			log.Println(obj) //这里可以切换到逻辑线程
		}
	}
}
