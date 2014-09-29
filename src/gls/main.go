// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"net"
	"gls/conf"
	"fmt"
	"time"
	"io"
	"encoding/binary"
)

func main() {
	defer func() {
		if x := recover(); x != nil {

		}
	}()

	go HandleSignal()
	go HandleGC()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", conf.GlsHost)
	if err != nil {
		panic(fmt.Sprintf("ResolveTCPAddr return error: %v", err))
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(fmt.Sprintf("ListenTCP return error: %v", err))
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {

	// create read channel
	header := make([]byte, 2)
	inBuffer := make(chan []byte, conf.InQueueSize)
	bufCtrl := make(chan bool)
	outBuffer := NewBuffer(conn, bufCtrl)

	defer func() {
		close(bufCtrl)
		close(inBuffer)
	}()

	// handle response to client
	go outBuffer.HandleClientRsp()

	for {
		// set tcp timeout
		conn.SetReadDeadline(time.Now().Add(TCP_READ_TIMEOUT * time.Second))

		// header
		n, err := io.ReadFull(conn, header)
		if err != nil {
			if err != io.EOF {
				// log
				println(n)
			}
			break
		}

		// data
		size := binary.BigEndian.Uint16(header)
		data := make([]byte, size)
		n, err = io.ReadFull(conn, data)
		if err != nil {
			if err != io.EOF {
				// log
			}
			break
		}

		select {
		case inBuffer <- data:
		case <- time.After(MAX_DELAY_IN * time.Second):
			return
		}
	}

}
