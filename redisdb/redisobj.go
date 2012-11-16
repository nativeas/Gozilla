package redisdb

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

/*[redisobj 基本结构]*/
type RedisObj struct {
	Nclient int
	c       redis.Conn
}

/*[创建RedisObj连接]*/
func CreateNewRedisConn(nID int, address string, password string) (*RedisObj, error) {
	var err error
	obj := new(RedisObj)
	obj.Nclient = nID
	obj.c, err = redis.Dial("tcp", address)
	if err != nil {
		log.Println("New RedisConn Error!")
		return nil, err
	}

	if _, err := obj.c.Do("AUTH", password); err != nil {
		log.Println("AUTH RedisConn Error!")
		obj.RedisObjClose()
		return nil, err
	}
	return obj, err
}

/*[关闭redis连接]*/
func (obj *RedisObj) RedisObjClose() {
	obj.c.Close()
}

/*[执行CMD至redis]*/
func (obj *RedisObj) Send(commandName string, args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(commandName, args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[从设定数据到Redis中]*/
func (obj *RedisObj) Set(args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(GetRedisCmd(CMD_SET), args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[从Redis中查找数据]*/
func (obj *RedisObj) Get(args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(GetRedisCmd(CMD_GET), args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[从设定数据到RedisHASH中]*/
func (obj *RedisObj) HSet(args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(GetRedisHashCmd(CMDHASH_SET), args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[从RedisHASH中获取数据]*/
func (obj *RedisObj) HGet(args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(GetRedisHashCmd(CMDHASH_GET), args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[批量设定数据到RedisHASH中]*/
func (obj *RedisObj) HMSet(args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(GetRedisHashCmd(CMDHASH_MSET), args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[批量从RedisHASH中获取数据]*/
func (obj *RedisObj) HMGet(args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(GetRedisHashCmd(CMDHASH_MGET), args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[转换成String]*/
func ConvertString(value interface{}) (reply string, err error) {
	var e0 error
	r, e := redis.String(value, e0)
	if e != nil {
		log.Println("ConverString Fail!")
	}
	return r, e
}
