package socket

import (
	"container/list"
	"log"
	"net"
	"os"
)

type RemoteRoom struct {
	listenAddr string
	remoteList *list.List
}

func NewRemoteRoom(listenAddr string) RemoteRoom {
	r := new(RemoteRoom)
	r.listenAddr = listenAddr
	r.remoteList = list.New()
	return *r
}

func (r *RemoteRoom) StartDaemon() {
	r.daemon()
}

func (r *RemoteRoom) daemon() {
	lis, err := net.Listen("tcp", r.listenAddr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer lis.Close()

	for {
		log.Println("waiting for new connection!")
		conn, error := lis.Accept()
		if error != nil {
			log.Println(error)
			os.Exit(1)
		}

		go r.connHandler(conn)
	}
}

func (r *RemoteRoom) connHandler(conn net.Conn) {
	log.Println("connhandler")
	remote := NewRemoteObject(conn)
	go remote.Read()
	r.remoteList.PushBack(remote)
	log.Println(r.remoteList.Len())
}
