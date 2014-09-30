// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"runtime"
	"time"
)

func HandleGC() {
	// timer
	gcTimer := time.NewTicker(time.Duration(GC_INTERVAL) * time.Second)

	for {
		select {
		case <-gcTimer.C:
			runtime.GC()
		}
	}
}
