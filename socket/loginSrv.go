package GoPray

import (
	"container/list"
	"fmt"
	"net"
	"os"
)

type LoginServer struct {
	listenAddr string
	gateList   *list.List
	logch chan string

}

func NewLoginServer(listenAddr string, ch chan string) LoginServer {
	ch <- "NewLoginServer"
	s := new(LoginServer)
	s.listenAddr = listenAddr
	s.gateList = list.New()
	s.logch = ch
	s.startDaemon()
	return *s
}

func (s *LoginServer) log(msg string){
	s.logch <-msg
}
func (s *LoginServer) startDaemon() {
	s.log("startDaemon")
	lis, error := net.Listen("tcp", s.listenAddr)
	if error != nil {
		fmt.Print("error")
		os.Exit(1)
	}

	defer lis.Close()

	for {
		fmt.Println("start accpet")
		conn, error := lis.Accept()
		if error != nil {
			fmt.Print("error")
			os.Exit(2)
		}

		go s.gateHandler(conn)
	}

}

func (s *LoginServer) gateHandler(conn net.Conn) {
	fmt.Println("ggateHandler()")
	gate := new(RemoteInstance)
	gate.Init(conn)
	go gate.Read()
	gate.Input <- []byte("Hello,Cclient")
	gate.Output <- []byte("dumyOutput")
	s.gateList.PushBack(gate)
}

