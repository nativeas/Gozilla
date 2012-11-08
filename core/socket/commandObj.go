package socket

import (
	"../packet"
)

type SocketCommand struct {
	packet.GozillaPacket
	Content string
}

type SocketPackage interface{}

func NewSocketCommand(m byte, s byte, c string) SocketPackage {
	obj := new(SocketCommand)
	obj.Content = c
	obj.MainCmd = m
	obj.SubCmd = s
	return *obj
}

//这个结构用来从服务里面输出到业务逻辑
//
type TargetdCommand struct {
	TaretId int           //目标对象id，用来对应remoteObj的id
	Command SocketPackage //发送的数据包
}
