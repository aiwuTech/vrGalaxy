// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import "center/log"

func main() {
	defer func() {
		if x := recover(); x != nil {
			log.E("center ")
		}
	}()
}

