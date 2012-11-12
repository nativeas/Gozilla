package socket

import (
	"../packet"
	"encoding/gob"
	// "errors"
	"log"
	"net"
	// "reflect"
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
	enc        *gob.Encoder
	dec        *gob.Decoder
}

func NewRemoteObject(conn net.Conn, id int) RemoteObject {
	log.Println("new remote object")
	object := new(RemoteObject)
	object.Conn = conn
	object.ObjId = id
	object.Input = make(chan packet.IGozillaPacket)
	object.Output = make(chan SocketCommand)
	object.Close = make(chan int)
	object.dec = gob.NewDecoder(object.Conn)
	object.enc = gob.NewEncoder(object.Conn)
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
	// obj:= new(gobObj)
	// // obj.Obj = cmd
	// err := r.enc.Encode(cmd)
	// if err != nil {
	// 	log.Println("SEND()")
	// 	log.Fatal(err)
	// }
	codec := new(PCodec)
	codec.Write(r.Conn, cmd)
	log.Println("send complete")
}

func (r *RemoteObject) Read() (cmd packet.IGozillaPacket, err error) {

	codec := new(PCodec)
	readObj, _ := codec.Read(r.Conn)
	// readedObj := new(interface{})
	// var q SocketCommand
	// errb := r.dec.Decode(&readedObj)
	// if errb != nil {
	// 	log.Println("READ()", errb)
	// 	//make some quit
	// 	return *readedObj, errors.New("client disconnect")
	// }
	// r.Output <- *readedObj
	return readObj, nil
}
