package main

import (
	"../core/player"
	"../core/socket"
	"./modules/login"
	"log"
)

// 实现了一个服务器的机制
// s.Output用来切换到逻辑线程
func main() {
	_ = login.NewPreLogin()
	var r player.IPlayerComamndRouter = new(player.PlayerCollection)
	s := socket.NewRemoteRoom("127.0.0.1:8080", r)
	log.Print(s)
	go s.StartDaemon()
	for {
		select {
		case obj := <-s.Output:
			log.Println(obj) //这里可以切换到逻辑线程
			log.Println("Push to Gate Logic to push to Server")

		}
	}
}
