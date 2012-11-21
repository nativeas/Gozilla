package datamodel

import (
	"../localdb"
	"../modeldb/playermodeldb"
	"../redisdb"
	"log"
	//	"reflect"
)

/*[记录连接redis对象用途小号]*/
const (
	REDISCONNECT_DATARES = 0
	REDISCONNECT_LOGRES  = 1
	REDISCONNECT_MAXID   = 2
)

type BaseModel struct {
	MdRedisCon  [REDISCONNECT_MAXID]*redisdb.RedisObj //redisdb 连接对象数组
	MdPlayerMap playermodeldb.PlayerObjMap            //玩家数据map
}

/*[初始化 BaseModel]*/
func InitModel() *BaseModel {
	log.Println("Begin Init BaseModel!")
	b := new(BaseModel)
	if b == nil {
		log.Println("Init BaseModel Fail!")
		return nil
	}
	var err error
	b.MdRedisCon[REDISCONNECT_DATARES], err = redisdb.CreateNewRedisConn(REDISCONNECT_DATARES, "127.0.0.1:6379", "1")
	if err != nil {
		log.Println("Redis Conn REDISCONNECT_DATARES Fail!")
		return nil
	}

	b.MdPlayerMap = playermodeldb.MakePlayerObjMap()
	return b
}

/*[关闭 redis连接]*/
func (mod *BaseModel) Close() {
	mod.MdRedisCon[REDISCONNECT_DATARES].RedisObjClose()
	mod.MdRedisCon[REDISCONNECT_LOGRES].RedisObjClose()
}

/*[获得 Table字段 对应数据 每增加一模块需增加一个分支 bo返回false表明无此数据]*/
func (mod *BaseModel) GetTableField(NTableID int) (reply []string, bo bool) {
	switch NTableID {
	case localdb.REDISDBINDEX_PLAYER:
		{
			return playermodeldb.PlayerObjField, true
		}
		break
	}
	return nil, false
}

/*[从BaseModel查找是否 有对应数据 每增加一模块需增加一个分支 bo返回false表明无此数据]*/
func (mod *BaseModel) FindTargetByTableIDHashkey(NHashTableID int, NHashKey int) (reply interface{}, bo bool) {
	var data interface{}
	success := false
	ok := false
	switch NHashTableID {
	case localdb.REDISDBINDEX_PLAYER:
		{
			data, ok = mod.MdPlayerMap.FindTargetByID(NHashKey)
			if ok {
				success = true
			}
		}
		break
	case localdb.REDISDBINDEX_ACCOUNT:
		{

		}
		break
	}
	return data, success
}

/*[更新数据到 远程redis]*/
func (mod *BaseModel) UpdateRedisDB(NConnID int, NHashTableID int, NHashKey int) bool {
	if NConnID < 0 || NConnID >= REDISCONNECT_MAXID {
		return false
	}

	if NHashTableID < 0 || NHashTableID >= localdb.REDISDBINDEX_MAX {
		return false
	}

	key, bo := localdb.GetRedisDBHashKey(0, "fengjie")
	if bo == false {
		log.Println("GetRedisDBHashKey Error!")
		return false
	}

	datastruct, bofind := mod.FindTargetByTableIDHashkey(NHashTableID, NHashKey)
	if bofind == false {
		log.Println("Not Exist Data In BaseModel!")
		return false
	}

	datafield, bofield := mod.GetTableField(NHashTableID)
	if bofield == false {
		log.Println("Get Table Field Fail!")
		return false
	}

	/*[更新 远端redis数据]*/
	p, _ := mod.MdRedisCon[NConnID].HMSet(localdb.TransHMSetDB(key, datafield, datastruct)[0:]...)
	log.Println(p)
	return true
}
