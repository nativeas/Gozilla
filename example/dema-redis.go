package main

import (
	"../localdb"
	"../modeldb/playermodeldb"
	"../redisdb"
	"fmt"
	"log"
	"strconv"
)

var s = []interface{}{
	"name", "fj007",
}

func main() {
	fmt.Println("Hello World")
	redisobj, err := redisdb.CreateNewRedisConn(0, "127.0.0.1:6379", "1")
	if err != nil {
		log.Println("Redis Conn Fail!")
		return
	}
	defer redisobj.RedisObjClose()

	/*[Test 设定值]*/
	//v, _ := redisobj.Set("name", "hy00")
	//log.Println(v)
	//var i = [10]interface{}{}
	//log.Println(len(i))
	var p playermodeldb.PlayerObjDB
	p.NOnlyID = 1
	p.StrCretName = "fengjie"

	var mm [3]string
	sss := 1
	mm[0] = strconv.Itoa(sss)
	mm[1] = "fff"
	mm[2] = string("333")
	log.Println(mm)

	/*[Test HMSET 标准参数转换数据]*/
	player, e := redisobj.HMSet(localdb.TransHMSetDB("fengjie", playermodeldb.PlayerObjField, p)[0:]...)
	if e != nil {
		log.Println(e)
	}
	log.Println(player)
	/*[Test HMGET 获取单个键值数据]*/
	playerdata, _ := redisobj.HMGet(localdb.TransHMGetDB("fengjie", playermodeldb.PlayerObjField)[0:]...)
	log.Println(playerdata)

	/*[Test 固定interface ]*/
	v, _ := redisobj.Set(s[0:]...)
	log.Println(v)

	/*[Test 获取值]*/
	v0, _ := redisobj.Get("name")
	log.Println(v0)

	/*[Test 转换String*/
	v1, _ := redisdb.ConvertString(v0)
	log.Println(v1)

	/*[Test 删除键值*/
	v2, _ := redisobj.Send(redisdb.GetRedisCmd(redisdb.CMD_DEL), "name")
	log.Println(v2)

	/*[Test Get hashkey]*/
	v3, bo := localdb.GetRedisDBHashKey(0, "fengjie")
	if bo != true {
		log.Println("GetRedisDBHashKey Error!")
	}
	log.Println(v3)
}
