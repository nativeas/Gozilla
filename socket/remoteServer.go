package socket

import (
	"container/list"
	"log"
	"net"
	"os"
)

type RemoteRoom struct {
	listenAddr   string
	remoteList   *list.List
	Output       chan TargetdCommand
	nextClientId int
}

func NewRemoteRoom(listenAddr string) RemoteRoom {
	r := new(RemoteRoom)
	r.listenAddr = listenAddr
	r.remoteList = list.New()
	r.Output = make(chan TargetdCommand)
	r.nextClientId = -1
	return *r
}

func (r *RemoteRoom) StartDaemon() {
	r.daemon()
}

func (r *RemoteRoom) daemon() {
	lis, err := net.Listen("tcp", r.listenAddr)
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	defer lis.Close()

	for {
		log.Println("waiting for new connection!")
		conn, error := lis.Accept()
		if error != nil {
			log.Println(error)
			os.Exit(2)
		}

		go r.connHandler(conn)
	}
}

func (r *RemoteRoom) connHandler(conn net.Conn) {
	log.Println("connhandler")
	r.nextClientId += 1
	remote := NewRemoteObject(conn, r.nextClientId)
	r.remoteList.PushBack(remote)
	go r.remoteObjectReader(remote)
	log.Println(r.remoteList.Len())
}

func (r *RemoteRoom) remoteObjectReader(obj RemoteObject) {
	for {
		cmd, err := obj.Read()
		if err != nil {
			log.Println(err)
			r.closeRemoteObject(obj)
			break
		}

		tcmd := new(TargetdCommand)
		tcmd.Command = cmd
		tcmd.TaretId = obj.ObjId
		// log.Printf("OUTPUT TCMD,target:%s,mcmd:%d,smcd:%d,content:%s",
		// 	tcmd.TaretId, cmd.MainCMD, cmd.SubCMD, cmd.ComandContent)
		r.Output <- *tcmd
	}
}

func (r *RemoteRoom) closeRemoteObject(obj RemoteObject) {
	//close
	obj.Close <- -1
	for item := r.remoteList.Front(); item != nil; item = item.Next() {
		remote := item.Value.(RemoteObject)
		if remote.RemoteAddr == obj.RemoteAddr {
			r.remoteList.Remove(item)
			break
		}
	}
}
