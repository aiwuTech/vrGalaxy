// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"gls/proto/api"
	"sync"
	"utils/packet"
	"gls/log"
)

var (
	wg  sync.WaitGroup
	die chan bool // for gls close
)

func init() {
	die = make(chan bool)
}

func HandleClientReq(in chan []byte, out *OutBuffer) {
	wg.Add(1)
	defer wg.Done()

	// main message loop
	for {
		select {
		case msg, ok := <-in:
			if !ok {
				return
			}

			if result := userReqProxy(msg); true {
				out.Send(result)
			}
		}
	}
}

func userReqProxy(p []byte) []byte {

	// read client request
	reader := packet.Reader(p)

	// read protocol number
	service, err := reader.ReadS16()
	if err != nil {
		log.E("gls read client request protocol return error: %v", err)
		return nil
	}
	serviceName, ok := api.ServiceList[service]
	if !ok {
		log.W("client request protocol[%v] not bind in gls server", service)
		return nil
	}

	log.T("client request protocol: %v", serviceName)

	// handle packet
	return api.SafeProtoHandler(service, reader)
}
