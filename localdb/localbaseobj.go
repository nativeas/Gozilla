package localdb

import (
	"../redisdb"
	"log"
	"reflect"
)

/*[转换 HASH 表设定]*/
func TransHMSetDB(hashname interface{}, field []string, i interface{}) []interface{} {
	v := reflect.ValueOf(i)
	if v.NumField() != len(field) {
		log.Println(v.Type(), "struct", v.NumField(), "!=", "field", len(field))
		return nil
	}

	db := make([]interface{}, v.NumField()+len(field)+1)
	db[0] = hashname //设定HASH 主键
	j := 1
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		db[j] = field[i]
		db[j+1] = value.Interface()
		j = j + 2
	}
	return db
}

/*[转换 HASH 读取]*/
func TransHMGetDB(hashname interface{}, field []string) []interface{} {
	length := len(field)
	if length < 1 {
		log.Println("Field less then zero!")
		return nil
	}
	db := make([]interface{}, length+1)
	db[0] = hashname //设定HASH 主键
	j := 1
	for i := 0; i < length; i++ {
		db[j] = field[i]
		j++
	}
	return db
}

/*[获得 HashKey index为db序号 i 为db主键]*/
func GetRedisDBHashKey(index int, i interface{}) (hashkey string, isconvert bool) {
	if index < 0 || index > REDISDBINDEX_MAX {
		log.Println("DBIndex out of Range!")
		return hashkey, false
	}
	key, e := redisdb.ConvertString(i)
	if e != nil {
		log.Println("ConvertString Interface Error!")
		return hashkey, false
	}
	hashkey = RedisDBHashKeyArray[index] + key
	return hashkey, true
}
