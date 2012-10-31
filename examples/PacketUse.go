package main

import (
	"../packet"
	"fmt"
	"log"
)

func main() {
	stlogin := new(packet.STPreLogin)
	stlogin.InitSubCmd()

	stloginret := new(packet.STPreLoginRet)
	stloginret.InitSubCmd()

	Process(stlogin)
	Process(stloginret)
	log.Println("Begin To Packet Process")
}

type I interface {
}

func Process(p I) {
	switch p.(type) {
	case *packet.STPreLogin:
		{
			log.Println("Recv 0x0101")
		}
		break
	case *packet.STPreLoginRet:
		{
			log.Println("Recv 0x0102")
		}
	default:
		log.Println("Packet Error")
	}
}
