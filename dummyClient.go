package GoPray

import (
	"net"
	"fmt"
	"os"
)

type user struct {
	name string
	pass string
}
type DummyClient struct {
	remoteAddr string
	User       user
	rinst	RemoteInstance
}

func Client() {
	fmt.Println("Hello,Client")
}

func (dc *DummyClient) init(raddr string, username string, password string) {
	dc.remoteAddr = raddr
	dc.User.name = username
	dc.User.pass = password
}

func (dc *DummyClient) login(){
	conn,err:= net.Dial("tcp",dc.remoteAddr)
	if err!=nil{
	   fmt.Println(err)
	   os.Exit(1)
	}
	ins:= new(RemoteInstance)
	dc.rinst =ins.Init(conn)
}

func (dc *DummyClient) authHandler(){

}
