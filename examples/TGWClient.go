package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//实现了一个狗日的发送给腾讯的包结构
func main() {
	msg := "tgw_l7_forward\r\nHost: s4.app100666618.qqopenapp.com:8001\r\n\r\n "
	// msg := "Hello"
	conn, err := net.Dial("tcp", "s4.app100666618.qqopenapp.com:8001")
	// conn, err := net.Dial("tcp", "172.16.2.56:8001")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	in, error := conn.Write([]byte(msg))
	if error != nil {
		fmt.Printf("Error sending data: %s, in: %d\n", error, in)
		os.Exit(2)
	}

	fmt.Println("Connection OK")
	data := make([]byte, 1024)
	n, error := conn.Read(data)
	if error != nil {
		log.Fatal(error)
	}
	log.Println(string(data[0:n]))
}
