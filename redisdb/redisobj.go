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

/*[设定数据至redis]*/
func (obj *RedisObj) Send(commandName string, args ...interface{}) (reply interface{}, err error) {
	obj.c.Send(commandName, args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[从Redis中查找数据]*/
func (obj *RedisObj) Get(args ...interface{}) (reply interface{}, err error) {
	obj.c.Send("Get", args...)
	obj.c.Flush()
	v, err := obj.c.Receive()
	return v, err
}

/*[转换成String]*/
func (obj *RedisObj) ConvertString(value interface{}) (reply string, err error) {
	var e0 error
	r, e := redis.String(value, e0)
	if e != nil {
		panic(e)
	}
	return r, e
}
