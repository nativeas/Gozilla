package main

import (
	"../engineer"
	//"../player"
	"../socket"
	"log"
)

func main() {
	s := socket.NewRemoteRoom("127.0.0.1:8080")
	log.Print(s)

	playerMap := engineer.CreateEngineer()
	player := playerMap.MakeNewPlayerObj(0)
	if player == nil {
		log.Println("Create player fail!")
	}

	cmd := socket.NewSocketCommand(1, 2, "")
	tarcmd := new(socket.TargetdCommand)
	tarcmd.Command = cmd
	tarcmd.TaretId = 0
	player.PutMsg(tarcmd)
	for {
		msg := player.GetMsg()
		if msg != nil {
			player.RunAll(*msg)
		}
	}

	go s.StartDaemon()
	for {
		select {
		case obj := <-s.Output:
			log.Println(obj) //这里可以切换到逻辑线程
			player.PutMsg(&obj)
		}

	}
}
