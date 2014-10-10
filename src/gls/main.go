// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"encoding/binary"
	"fmt"
	"gls/conf"
	"io"
	"net"
	"time"
	"gls/log"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			log.E("gls caught panic: %v", x)
		}
	}()

	go HandleSignal()
	go HandleGC()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", conf.GlsHost)
	if err != nil {
		panic(fmt.Sprintf("ResolveTCPAddr return error: %v", err))
	}

	log.I("trying to startup gls server: %v", tcpAddr)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		panic(fmt.Sprintf("ListenTCP return error: %v", err))
	}
	defer listener.Close()

	log.I("gls server started, listening...")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.T("gls accept client connection return error: %v", err)
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
		conn.Close()
	}()

	// handle response to client
	go outBuffer.HandleClientRsp()
	// handle request from client
	go HandleClientReq(inBuffer, outBuffer)

	for {
		// set tcp timeout
		conn.SetReadDeadline(time.Now().Add(TCP_READ_TIMEOUT * time.Second))

		// header
		n, err := io.ReadFull(conn, header)
		if err != nil {
			if err != io.EOF {
				// log
				log.T("gls read client request header return error: %v, read: %v", err, n)
			}
			break
		}
		log.T("header: %v", header)

		// data
		size := binary.BigEndian.Uint16(header)
		log.T("size: %v", size)
		data := make([]byte, size)
		n, err = io.ReadFull(conn, data)
		if err != nil {
			if err != io.EOF {
				// log
				log.T("gls read client request body return error: %v, read: %v", err, n)
			}
			break
		}

		select {
		case inBuffer <- data:
		case <-time.After(MAX_DELAY_IN * time.Second):
			return
		}
	}

}
