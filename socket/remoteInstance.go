package GoPray

import (
	"fmt"
	"log"
	"net"
)

//
// Impl a remote Instance for client 
//
type RemoteInstance struct {
	Id         int
	ListenAddr string
	Conn       net.Conn
	Input      chan []byte
	Output     chan []byte
	Dispose    chan int
}

//
// RemoteInstance Init
//
func (r *RemoteInstance) Init(conn net.Conn) {
	r.Conn = conn
	r.ListenAddr = conn.RemoteAddr().String()
	r.Input = make(chan []byte)
	r.Output = make(chan []byte)
	go r.daemon()
}

func (r *RemoteInstance) send(msg []byte) {
	r.Conn.Write(msg)
}

func (r *RemoteInstance) daemon() {
	for {
		select {
		case buffer := <-r.Input:
			fmt.Println("do input ")
			r.send(buffer)
		case out := <-r.Output:
			log.Print(string(out))
		case <-r.Dispose:
			break
		}

	}
}

func (r *RemoteInstance) Read() {
	readBuffer := make([]byte, 2048)
	_, error := r.Conn.Read(readBuffer)
	if error != nil {
		fmt.Print("Client Error!")
		fmt.Print(error)
		r.Dispose <- 1
	} else {
		r.Output <- readBuffer
		r.Read()
	}

}
