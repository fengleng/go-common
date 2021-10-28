//go:build windows
// +build windows

package atexit

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Cb func()

var (
	signalChan       = make(chan os.Signal, 1)
	exitCallbackList []Cb
	cbLock           sync.Mutex
)

func init() {
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGILL,
		syscall.SIGABRT,
		syscall.SIGBUS,
		syscall.SIGFPE,
		syscall.SIGSEGV,
		//syscall.SIGPIPE,
	)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				signalChan <- syscall.SIGINT
			}
		}()
		<-signalChan
		cbLock.Lock()
		for _, cb := range exitCallbackList {
			cb()
		}
		cbLock.Unlock()
	}()
}

func Register(cbList ...Cb) {
	cbLock.Lock()
	exitCallbackList = append(exitCallbackList, cbList...)
	defer cbLock.Unlock()
}
