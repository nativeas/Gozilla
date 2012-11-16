package redisdb

import (
	"log"
)

/*[CMD数组序号]*/
const (
	CMD_GET    = 0 //根据Key获取Value
	CMD_SET    = 1 //设定Key&Value
	CMD_DEL    = 2 //根据key删除键值对
	CMD_EXISTS = 3 //根据Key判断是否存在
	CMD_INCR   = 4 //根据Key自增Value(只适好用于 integer)
	CMD_APPEND = 5 //根据Key增加Value值
	CMD_MAX    = 6 //最大下标索引边界
)

/*[CMD数组]*/
var CommandArray = []string{
	"GET",
	"SET",
	"DEL",
	"EXISTS",
	"INCR",
	"APPEND",
}

/*[超出界限 默认为选择GetCmd]*/
func GetRedisCmd(nCmdID int) string {
	if nCmdID < 0 || nCmdID >= CMD_MAX {
		log.Printf("CmdID %d CMD out!", nCmdID)
		return CommandArray[CMD_GET]
	}
	return CommandArray[nCmdID]
}

/*[HASHCMD数组序号]*/
const (
	CMDHASH_GET    = 0 //根据Key获取Value
	CMDHASH_SET    = 1 //设定Key&Value
	CMDHASH_DEL    = 2 //根据key删除键值对
	CMDHASH_EXISTS = 3 //根据Key判断是否存在
	CMDHASH_INCRBY = 4 //根据Key自增Value(只适好用于 integer 可以定义增加数量)
	CMDHASH_LEN    = 5 //获取HASH成员个数
	CMDHASH_MSET   = 6 //批量设定HASH成员数据
	CMDHASH_MGET   = 7 //批量获得HASH成员数据
	CMDHASH_MAX    = 8 //最大下标索引边界
)

/*[HASHCMD数组]*/
var HashCommandArray = []string{
	"HGET",
	"HSET",
	"HDEL",
	"HEXISTS",
	"HINCRBY",
	"HLEN",
	"HMSET",
	"HMGET",
}

/*[超出界限 默认为选择HGetCmd]*/
func GetRedisHashCmd(nCmdID int) string {
	if nCmdID < 0 || nCmdID >= CMDHASH_MAX {
		log.Printf("CmdID %d CMDHASH out!", nCmdID)
		return HashCommandArray[CMDHASH_GET]
	}
	return HashCommandArray[nCmdID]
}
