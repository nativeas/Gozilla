package engine

import (
	"../packet"
	"container/list"
)

type PlayerObj struct {
	NclientId int
	Lmsg      *list.List
}

func newPlayerObj(nid int) *PlayerObj {
	obj := new(PlayerObj)
	obj.NclientId = nid
	obj.Lmsg = list.New()
	return obj
}

type PlayerCollection struct {
	players map[int]*PlayerObj
}

func (p *PlayerCollection) Init() {
	p.players = make(map[int]*PlayerObj)
}

func (p *PlayerCollection) CreatePlayer(Nclient int) {
	p.players[Nclient] = newPlayerObj(Nclient)
}

func (p *PlayerCollection) RemovePlayer(Nclient int) {
	delete(p.players, Nclient)
}

func (p *PlayerCollection) Search(NclientID int) *PlayerObj {
	player, ok := p.players[NclientID]
	if ok {
		return player
	}
	return nil
}

func (p *PlayerCollection) Pump() (*PlayerObj, packet.IGozillaPacket) {
	return nil, nil
}

func (p *PlayerCollection) PushPacket(NclientId int, packet packet.IGozillaPacket) {

}
