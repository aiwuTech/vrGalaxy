package main

import (
	"os"
	"os/signal"
	"syscall"
)

//----------------------------------------------- handle unix signals
func HandleSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM)

	for {
		msg := <-ch

		switch msg {
		case syscall.SIGHUP:
		case syscall.SIGTERM:
		}
	}
}
