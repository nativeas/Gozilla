package main

import (
	"../core/engine"
	// "../core/player"
	"../core/socket"
	// "./modules/login"
	"log"
)

// 实现了一个服务器的机制
// s.Output用来切换到逻辑线程
func main() {
	eng := new(engine.Engine)
	eng.Init("alpha1")
	s := socket.NewRemoteRoom("127.0.0.1:8080", eng)
	log.Print(s)
	s.StartDaemon()
}
