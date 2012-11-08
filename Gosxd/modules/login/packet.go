package login

import (
	"../../../core/packet"
	"../../modules"
)

const (
	CMD_PRELOGIN       = 1 //c->s
	CMD_PRELOGIN_REPLY = 2 //s->c
	CMD_LOGIN          = 3 //c->s
	CMD_LOGIN_REPLY    = 4 //s->c
	CMD_LOGIN_TOKEN    = 5 //c->s
)

type basePacket struct {
	packet.GozillaPacket
}

func (l *basePacket) Init() {
	l.MainCmd = modules.MOD_LOGIN
}

type PreLogin struct {
	basePacket
}

func (p *PreLogin) Init() {
	p.basePacket.Init()
	p.SubCmd = CMD_PRELOGIN
}

type PreLogin_Reply struct{}{
	basePacket
}

