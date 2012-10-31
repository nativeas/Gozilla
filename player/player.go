package player

import (
	"../socket"
	"container/list"
	"log"
)

/*人物结构*/
type PlayerObj struct {
	NclientID int
	Lmsg      *list.List
}

/*创建新玩家*/
func CreateNewPlayer(NclientID int) *PlayerObj {
	p := new(PlayerObj)
	p.NclientID = NclientID
	p.Lmsg = list.New()
	return p
}

/*获取数据包 每次取一个*/
func (p *PlayerObj) GetMsg() (m *socket.TargetdCommand) {
	for msg := p.Lmsg.Front(); msg != nil; msg = msg.Next() {
		if msg != nil {
			msgconvert := msg.Value.(socket.TargetdCommand)
			p.Lmsg.Remove(msg)
			return &msgconvert
		}
	}
	return nil
}

/*将消息压入队列*/
func (p *PlayerObj) PutMsg(msg *socket.TargetdCommand) {
	if msg != nil {
		p.Lmsg.PushBack(*msg)
	}
}

/*总逻辑入口*/
func (p *PlayerObj) RunAll(msg socket.TargetdCommand) {
	switch msg.Command.MainCMD {
	case 0x01: //登录模块
		log.Println("login")
		p.LoginProcess(msg)
		break
	}
}

/*登录模块逻辑入口*/
func (p *PlayerObj) LoginProcess(msg socket.TargetdCommand) {
	switch msg.Command.SubCMD {
	case 0x02:
		log.Println("loginret")
		break
	}
}
