package atexit

import (
	"log"
	"testing"
	"time"
)

func TestExit(t *testing.T) {
	Register(func() {
		log.Fatalln("hahhahha")
	})
	for {
		time.Sleep(2 * time.Second)
		t.Log("ggggg")

	}
}
