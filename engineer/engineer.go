package engineer

import (
	"../player"
	"log"
)

type PlayerObjMap struct {
	MplayerMap map[int]*player.PlayerObj
}

func CreateEngineer() *PlayerObjMap {
	p := new(PlayerObjMap)
	p.MplayerMap = make(map[int]*player.PlayerObj)
	return p
}

func (p *PlayerObjMap) MakeNewPlayerObj(NclientID int) *player.PlayerObj {
	playerobj := p.FindPLayerByID(NclientID)
	if playerobj != nil {
		return playerobj
	}
	/*创建新玩家*/
	player := player.CreateNewPlayer(NclientID)
	if player != nil {
		p.MplayerMap[NclientID] = player
		return player
	}
	return nil
}

func (p *PlayerObjMap) RemovePlayer(NclientID int) {
	delete(p.MplayerMap, NclientID)
}

func (p *PlayerObjMap) FindPLayerByID(NclientID int) *player.PlayerObj {
	player, ok := p.MplayerMap[NclientID]
	if ok {
		log.Println("Player has exist!")
		return player
	}
	return nil
}
