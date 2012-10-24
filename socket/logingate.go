package GoPray

import (
	"log"
	"net"
	"os"
)

func LoginGate() {
	startService()
}

/*
start gate service 
1. connect to server
2. waiting for client
*/
func startService() {
	log.Print("Start Service")

	serviceAddr := "127.0.0.1:8001"

	lis, error := net.Listen("tcp", serviceAddr)
	defer lis.Close()

	if error != nil {
		log.Fatal("Error!")
		os.Exit(1)
	}

	cm := make(map[string]net.Conn)
	for {
		log.Print(len(cm))
		conn, error := lis.Accept()
		if error == nil {
			cm[conn.RemoteAddr().String()] = conn
			go handlerConn(conn, cm)
		}
	}
}


func handlerConn(conn net.Conn, cm map[string]net.Conn) {
	log.Print("Connection Detected ,start handler")
	log.Print(conn.RemoteAddr())
	hello := "Hello,dude"
	conn.Write([]byte(hello))
	data := make([]byte, 1024)
	_, error := conn.Read(data)
	log.Print(string(data))
	if error != nil {
		log.Print("Client lost!")

		delete(cm, conn.RemoteAddr().String())

	}
}

/*
connect to login server

*/
func connectLoginServer() {
	loingSvrAddr := "127.0.0.1:8100"
	conn, error := net.Dial("tcp", loingSvrAddr)

	if error != nil {
		log.Fatal("fuck!")
		os.Exit(1)
	}

	for {
		b := make([]byte, 1024)
		_, error := conn.Read(b)
		if error != nil {
			log.Fatal("error")
			os.Exit(2)
		}
	}
}
