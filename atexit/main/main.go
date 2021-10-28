package main

import (
	"github.com/fengleng/go-common/atexit"
	"log"
	"time"
)

func main() {
	atexit.Register(func() {
		log.Fatalln("hahhahha")
	})
	for {
		log.Println("lalalla")
		time.Sleep(2 * time.Second)
		panic("fdsfasf")
	}
}
