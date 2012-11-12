package engine

import (
	"../packet"
	"container/list"
)

type PlayerObj struct {
	NclientId int
	Lmsg      *list.List
}

func newPlayerObj(nid int) *PlayerObj{
	obj:= new(PlayerObj)
	obj.NclientId = nid
	obj.Lmsg = list.New()
	return obj
}

func (p *PlayerObj) pushMsg(msg packet.IGozillaPacket) {
	p.Lmsg.PushBack(msg)
}

func (p *PlayerObj) getMsg(packet packet.IGozillaPakcet) {
	if p.Lmsg.Len() > 0 {
		p.Lmsg.PushBack(packet)
	}
}

type PlayerCollection struct {
	players map[int]PlayerObj
}

func (p *PlayerCollection) Init() {
	p.players = make(map[int]PlayerObj)
}

func (p *PlayerCollection) CreatePlayer(Nclient int) {
	map[Nclient] = newPlayerObj(Nclient)
}

func (p *PlayerCollection) RemovePlayer(Nclient int) {
	delete(p.players, Nclient)
}

func (p *PlayerCollection) Search(NclientID int) *PlayerObj {
	player,ok:=p.players[NclientID]
	if ok{
		return player
	}
	return nil
}


func (p *PlayerCollection) Pump() (p *PlayerObj, cmd packet.IGozillaPacket){
	return nil, nil
}