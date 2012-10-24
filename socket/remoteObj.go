package socket

import (
	"encoding/gob"
	"log"
	"net"
)

/*
远端对象
*/
type RemoteObject struct {
	RemoteAddr net.Addr
	Conn       net.Conn
	Input      chan SocketCommand
	Output     chan SocketCommand
	Close      chan int
	enc        *gob.Encoder
	dec        *gob.Decoder
}

func NewRemoteObject(conn net.Conn) RemoteObject {
	log.Println("new remote object")
	object := new(RemoteObject)
	object.Conn = conn
	object.Input = make(chan SocketCommand)
	object.Output = make(chan SocketCommand)
	object.Close = make(chan int)
	object.dec = gob.NewDecoder(object.Conn)
	object.enc = gob.NewEncoder(object.Conn)
	go object.daemon()
	return *object
}

func (r *RemoteObject) daemon() {
	for {
		select {
		case <-r.Close:
			log.Println("CLOSE")
			break
		case obj := <-r.Input:
			log.Println("INPUT")
			r.send(obj)
		case obj := <-r.Output:
			log.Printf("OUTPUT,mcmd:%d,smcd:%d,content:%s", obj.MainCMD, obj.SubCMD,
				obj.ComandContent)
		}
	}
}

func (r *RemoteObject) send(cmd SocketCommand) {
	err := r.enc.Encode(cmd)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send complete")
}

func (r *RemoteObject) Read() {
	readedObj := new(SocketCommand)
	// var q SocketCommand
	err := r.dec.Decode(&readedObj)
	if err != nil {
		log.Fatal(err)
	}
	r.Output <- *readedObj
	r.Read()
}
