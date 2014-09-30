// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"code.google.com/p/goprotobuf/proto"
	"gls/proto/pb"
	"net"
	"utils/packet"
	"time"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8888")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	reqMsg := new(pb.ReqHeartBeat)

	data, _ := proto.Marshal(reqMsg)

	cnt := 1
	for {

		println("sending heartbeat to server...", cnt)

		writer := packet.Writer()
		writer.WriteU16(uint16(len(data) + 4))
		writer.WriteS16(int16(0))
		writer.WriteBytes(data)

		conn.Write(writer.Data())

		time.Sleep(10 * time.Second)
		cnt ++
	}
}
