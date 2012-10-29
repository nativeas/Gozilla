package main

import (
	"../packet"
	"../test"
	"fmt"
)

func main() {
	test.Helloworld()
	s := new(packet.STPreLogin)
	s.InitSubCmd()
	fmt.Println(s)
	fmt.Printf("Begin To Program\n")
}
