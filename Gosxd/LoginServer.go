package main

import (
	"../core/engine"
	"../core/socket"
	"./modules/login"
	"log"
)

// LoginServer
func main() {
	lm := new(login.Login_Module)
	eng := new(engine.Engine)
	eng.Init()
	eng.RegisterModule(lm)
	s := socket.NewRemoteRoom("127.0.0.1:8001", eng)
	// s := socket.NewGateObjet("127.0.0.1:8080", "127.0.0.1:8001")
	log.Print(s)
	s.StartDaemon()
}
