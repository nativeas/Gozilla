package main

import (
	"../Gosxd/modules/login"
	"../core/packet"
	// "../core/socket"
	// "bytes"
	"log"
	"reflect"
)

func main() {

	var obj interface{} = new(login.PreLogin)
	var obt = obj.(packet.IGozillaPacket)
	log.Println(obj, obt)
	var typ = reflect.TypeOf(obj)
	log.Println(typ)
	var obj2 = reflect.New(typ).Kind()
	var v2 = reflect.ValueOf(obj)
	log.Println(obj2, v2.Kind())
	// log.Println(v2)
	// reflect.NewAt(typ, p)
	// buf := bytes.NewBuffer(nil)
	// codec := new(socket.PCodec)
	// cmd := new(login.PreLogin)
	// cmd.Init()
	// //cmd.(packet.IGozillaPacket)
	// codec.Write(buf, cmd)
	// socket.RegisterCommand(cmd)
	// obj, err := codec.Read(buf)
	// log.Println(obj, err)
}
