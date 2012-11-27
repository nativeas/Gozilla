package login

import (
	"../../../core/packet"
	"../../../core/socket"
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
	Name string
}

type UserLogin struct {
	basePacket
	Name     string
	Password string
}

func (p *PreLogin) Init() {
	p.basePacket.Init()
	p.SubCmd = CMD_PRELOGIN
}

func (u *UserLogin) Init() {
	u.basePacket.Init()
	u.SubCmd = CMD_LOGIN
	u.Name = "dan"
	u.Password = "123456"
}

func NewPreLogin() packet.IGozillaPacket {
	p := new(PreLogin)
	p.Init()
	return p
}

func NewUserLogin() packet.IGozillaPacket {
	p := new(UserLogin)
	p.Init()
	return p
}

func init() {
	socket.RegisterCommand(packet.IGozillaPacket(NewPreLogin()), NewPreLogin)
	socket.RegisterCommand(packet.IGozillaPacket(NewUserLogin()), NewUserLogin)
}
