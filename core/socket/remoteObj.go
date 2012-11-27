package socket

import (
	"../packet"
	"log"
	"net"
)

type cmdObj interface{}
type gobObj struct {
	Obj cmdObj
}

/*
远端对象
*/
type RemoteObject struct {
	ObjId      int
	RemoteAddr net.Addr
	Conn       net.Conn
	Input      chan packet.IGozillaPacket
	Output     chan SocketCommand
	Close      chan int
	codec      RemoteCodec
}

func NewRemoteObject(conn net.Conn, id int) RemoteObject {
	log.Println("new remote object")
	object := new(RemoteObject)
	object.Conn = conn
	object.ObjId = id
	object.Input = make(chan packet.IGozillaPacket)
	object.Output = make(chan SocketCommand)
	object.Close = make(chan int)
	object.codec = new(GobCodec2)
	go object.daemon()
	return *object
}

func (r *RemoteObject) daemon() {
DAEMON_LOOP:
	for {
		select {
		case <-r.Close:
			log.Println("CLOSE")
			r.Conn.Close()
			break DAEMON_LOOP
		case obj := <-r.Input:
			log.Println("INPUT")
			r.send(obj)
		}
	}
	log.Println("end deamon()")
}

func (r *RemoteObject) send(cmd packet.IGozillaPacket) {
	r.codec.Write(r.Conn, cmd)
	log.Println("send complete")
}

func (r *RemoteObject) Read() (cmd packet.IGozillaPacket, err error) {
	readObj, err := r.codec.Read(r.Conn)

	return readObj, err
}
