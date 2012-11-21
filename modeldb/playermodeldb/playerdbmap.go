package playermodeldb

import (
	"../../localdb"
	"log"
)

/*[玩家 redisdeb 管理 Map]*/
type PlayerObjMap struct {
	NIndex      int    //map 序号
	StrHashName string //map hashkey
	PlayerMap   map[int]PlayerObjDB
}

/*[初始化 函数]*/
func MakePlayerObjMap() PlayerObjMap {
	pmap := new(PlayerObjMap)
	pmap.PlayerMap = make(map[int]PlayerObjDB)
	pmap.NIndex = localdb.REDISDBINDEX_PLAYER
	pmap.StrHashName = localdb.RedisDBHashKeyArray[pmap.NIndex]
	return *pmap
}

/*[通过onlyID 查找数据]*/
func (pMap *PlayerObjMap) FindTargetByID(NOnlyID int) (p PlayerObjDB, bo bool) {
	p, ok := pMap.PlayerMap[NOnlyID]
	return p, ok
}

/*[map 增加对象]*/
func (pMap *PlayerObjMap) Add(NOnlyID int, playdb PlayerObjDB) bool {
	_, ok := pMap.FindTargetByID(NOnlyID)
	if ok {
		log.Println("NOnlyID", NOnlyID, "Exists!")
		return false
	}
	pMap.PlayerMap[NOnlyID] = playdb
	return true
}

/*[map 删除对象]*/
func (pMap *PlayerObjMap) Del(NonlyID int) {
	delete(pMap.PlayerMap, NonlyID)
}
