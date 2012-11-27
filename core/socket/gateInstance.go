package socket

import (
	"../packet"
	"log"
	"net"
)

type GateInstance struct {
	serverAddr     string //server address
	serverObject   RemoteObject
	listenAddr     string            //listened address
	listenRoom     RemoteRoom        //client list
	clientMsgQueue *gateMessageQueue //客户端消息队列
	serverMsgQueue *gateMessageQueue //服务端消息队列
}

func (g *GateInstance) PushPacket(NclientId int, packet packet.IGozillaPacket) {
	g.clientMsgQueue.PushPacket(NclientId, packet)
}

func NewGateObjet(listen string, server string) *GateInstance {
	obj := new(GateInstance)
	obj.listenAddr = listen
	obj.serverAddr = server
	obj.clientMsgQueue = newGateMessageQueue()
	obj.serverMsgQueue = newGateMessageQueue()
	obj.listenRoom = NewRemoteRoom(obj.listenAddr, obj.clientMsgQueue)

	return obj
}

func (g *GateInstance) StartDaemon() {
	go g.listenRoom.StartDaemon()
}

func (g *GateInstance) Dial() {
	g.connectServer()
}

func (g *GateInstance) connectServer() {
	conn, error := net.Dial("tcp", g.serverAddr)
	if error != nil {
		log.Fatal(error)
	}
	g.serverObject = NewRemoteObject(conn, 1)
	log.Println("server connected!")
}
