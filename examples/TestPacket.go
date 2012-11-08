package main

import (
	"../core/packet"
	"log"
)

//login packet sample

type LoginPacket struct {
	packet.GozillaPacket
}

const (
	LOGIN_MOD    = 1
	PRELOGIN_CMD = 1
)

func (l *LoginPacket) Init() {
	l.MainCmd = LOGIN_MOD
}

type PreloginPacket struct {
	LoginPacket
}

func (p *PreloginPacket) Init() {
	p.LoginPacket.Init()
	p.SubCmd = PRELOGIN_CMD
}

func main() {
	obj := new(PreloginPacket)
	obj.Init()
	log.Println(obj)
}
