package socket

import (
	"../packet"
	"log"
	"strconv"
)

var cmds map[string]func() packet.IGozillaPacket

func init() {
	cmds = make(map[string]func() packet.IGozillaPacket)
}

func RegisterCommand(cmd packet.IGozillaPacket, fun func() packet.IGozillaPacket) {
	k := GetCommandMark(cmd)
	log.Println("k", k)
	cmds[k] = fun
}

func GetCommand(t string) interface{} {
	if fun, ok := cmds[t]; ok {
		return fun()
	}
	log.Println(t, "is not exist in CommandMapping")
	return nil
}

func GetCommandMark(pkt packet.IGozillaPacket) string {
	mcmd := strconv.Itoa(int(pkt.GetMainCmd()))
	scmd := strconv.Itoa(int(pkt.GetSubCmd()))
	mark := "cmd@" + mcmd + "_" + scmd
	return mark
}
