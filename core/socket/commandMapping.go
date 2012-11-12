package socket

import (
	"../packet"
	"log"
	"reflect"
)

var cmds map[string]func() packet.IGozillaPacket

func init() {
	cmds = make(map[string]func() packet.IGozillaPacket)
}

func RegisterCommand(cmd packet.IGozillaPacket, fun func() packet.IGozillaPacket) {
	t := reflect.TypeOf(cmd)
	k := t.Name()
	cmds[k] = fun
}

func GetCommand(t string) packet.IGozillaPacket {
	if fun, ok := cmds[t]; ok {
		return fun()
	}
	log.Fatal(
		t, "is not exist in CommandMapping")
	return nil
}
