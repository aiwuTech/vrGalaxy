// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"sync"
	"utils/packet"
)

var (
	wg sync.WaitGroup
	die chan bool   // for gls close
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
		case msg, ok := <- in:
			if !ok {
				return
			}


		}
	}


}

func userReqProxy(p []byte) []byte {

	// read client request
	reader := packet.Reader(p)

	// read protocol number
	serverId, err := reader.ReadS16()
	if err != nil {
		return nil
	}



}
