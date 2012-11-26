package main

import (
	"../core/packet"
	"../core/socket"
	"./modules/login"
	"bytes"
	"log"
)

func main() {

	log.Println("start test")

	buf := bytes.NewBuffer(nil)
	codec := new(socket.GobCodec)
	var obj packet.IGozillaPacket = new(login.PreLogin)
	log.Println("to write")
	codec.Write(buf, obj)
	log.Println("write complte")
	for i := 0; i < 3; i++ {
		to, te := codec.Read(buf)
		log.Println("read complte")
		log.Println(to, te)
	}

}
