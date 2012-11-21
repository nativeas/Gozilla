package main

import (
	"../datamodel"
	"../modeldb/playermodeldb"
	"log"
)

func main() {
	/*[初始化 datamodel]*/
	mod := datamodel.InitModel()

	p := new(playermodeldb.PlayerObjDB)
	p.NOnlyID = 1
	p.StrCretName = "fj"

	/*[Test 添加数据到datamodel playermap]*/
	boAdd := mod.MdPlayerMap.Add(1, *p)
	if boAdd == false {
		log.Println("Add Fail!")
	}

	/*[Test 从playermap里查找对应数据]*/
	f, ok := mod.MdPlayerMap.FindTargetByID(1)
	if ok {
		log.Println("Exist!  FindTargetByID")
	}
	log.Println(f)

	/*[Test 从playermap中删除对象]*/
	//mod.MdPlayerMap.Del(0)

	/*[Test 更新数据到远端]*/
	mod.UpdateRedisDB(0, 0, 1)
}
