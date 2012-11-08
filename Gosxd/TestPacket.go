package main

import (
	"./modules/login"
	"log"
)

func main() {
	obj := new(login.PreLogin_Reply)
	obj.Init()
	log.Println(obj)
}
