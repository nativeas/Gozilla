package GoPray

import (
	"log"
	"net"
)

/*
Crossdomain() function impl the flash socket connection policy service.
*/
func Crossdomain() {
	courent := "<cross-domain-policy><allow-access-from domain=\"*\" to-ports=\"*\" /></cross-domain-policy>\n"
	ln, error := net.Listen("tcp", ":843")
	if error != nil {
		panic(error)
	}

	for {
		conn, error := ln.Accept()
		if error != nil {
			log.Fatal("get client connection Error: ", error)

		}

		data := []byte(courent)

		log.Println("Someone connected!")
		conn.Write(data)
		log.Print(conn.RemoteAddr(), conn.LocalAddr())
		conn.Close()
	}

}
