// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"net"
	"gls/conf"
	"utils/packet"
)

type OutBuffer struct {
	ctrl    chan bool   // receive exit signal
	pending chan []byte // pending packet
	max     int         // max queue size
	conn    net.Conn    // connection
}

func NewBuffer(conn net.Conn, ctrl chan bool) *OutBuffer {
	return &OutBuffer{
		ctrl: ctrl,
		pending: make(chan []byte, conf.OutQueueSize),
		max: conf.OutQueueSize,
		conn: conn,
	}
}

func (buf *OutBuffer) HandleClientRsp() {
	for{
		select {
		case data := <- buf.pending:
			buf.rawSend(data)
		case <- buf.ctrl:
			close(buf.pending)
			for data := range buf.pending {
				buf.rawSend(data)
			}
			buf.conn.Close()
			return
		}
	}
}

// send packet
func (buf *OutBuffer) Send(data []byte) {
	buf.pending <- data
}

func (buf *OutBuffer) rawSend(data []byte) {
	writter := packet.Writer()
	writter.WriteU16(uint16(len(data)))
	writter.WriteRawBytes(data)

	n, err := buf.conn.Write(writter.Data())
	if err != nil {
		println(n)
	}
}
