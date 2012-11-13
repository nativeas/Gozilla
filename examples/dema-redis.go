package main

import (
	"../redisdb"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World")
	redisobj, err := redisdb.CreateNewRedisConn(0, "127.0.0.1:6379", "1")
	if err != nil {
		log.Println("Redis Conn Fail!")
		return
	}
	defer redisobj.RedisObjClose()
	/*[Test 设定值]*/
	v, _ := redisobj.Send("Set", "name", "hy00")
	log.Println(v)

	/*[Test 获取值]*/
	v0, _ := redisobj.Get("name")
	log.Println(v0)

	/*[Test 转换String*/
	v1, _ := redisobj.ConvertString(v0)
	log.Println(v1)
}
