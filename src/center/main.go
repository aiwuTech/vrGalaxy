// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"center/log"
	"fmt"
	"net"
	"center/conf"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			log.E("center caught panic: %v", x)
		}
	}()

	go HandleSignal()
	go HandleGC()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", conf.CenterHost)
	if err != nil {
		panic(fmt.Sprintf("ResolveTCPAddr return error: %v", err))
	}

	log.I("trying to startup center server: %v", tcpAddr)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		panic(fmt.Sprintf("ListenTCP return error: %v", err))
	}

	log.I("center server started, listening...")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.T("center server accept connection return error: %v", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {

}
