package main

import (
	"../core/engine"
	"../core/socket"
	"log"
)

// 实现了一个服务器的机制
// eng用来接受所有的输出
// LoginGate 
func main() {
	eng := new(engine.Engine)
	eng.Init()
	s := socket.NewRemoteRoom("127.0.0.1:8080", eng)
	//s := socket.NewGateObjet("127.0.0.1:8080", "127.0.0.1:8001")
	log.Print(s)
	// s.Dial()
	s.StartDaemon()
	var ch chan int
	for {
		select {
		case <-ch:
			log.Print("He")
		}
	}
}
