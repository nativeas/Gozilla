package main

import (
	"container/list"
	"log"
	"time"
)

func main() {
	log.Println("HElLo")
	lt := list.New()

	t := time.Now()
	for i := 0; i < 100000; i++ {
		lt.PushBack(1)
	}

	log.Println(lt.Len())
	t2 := time.Now()
	for i := 0; i < 100000; i++ {
		v := lt.Front()
		lt.Remove(v)
	}

	log.Println(lt.Len())

	t3 := time.Now()

	log.Println("first time ", t2.Sub(t))
	log.Println("2nd time ", t3.Sub(t2))
	log.Println("total time ", t3.Sub(t))
}
