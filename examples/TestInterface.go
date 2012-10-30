package main

import (
	"log"
)

type doCall interface {
	call()
}

type duck struct {
	name string
}

func (d *duck) call() {
	log.Println("duck call!")
}

type fatduck struct {
	duck
}

func (f *fatduck) call() {
	log.Println("fatduck call!")
}

func main() {
	var d doCall = new(fatduck)
	d.call()
}
