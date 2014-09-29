// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package conf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

const (
	_GLS_SECTION      = "gls"
	_LOG_SECTION      = "log"
	_DEFAULT_GLS_HOST = ":8888"
	_DEFAULT_IN_SIZE  = 64
	_DEFAULT_OUT_SIZE = 15
)

var (
	confFilePath = "conf.ini"
	confObj      *goconfig.ConfigFile
)

// configure object
var (
	GlsHost      string
	InQueueSize  int
	OutQueueSize int
)

func init() {
	var err error
	confObj, err = goconfig.LoadConfigFile(confFilePath)
	if err != nil {
		panic(fmt.Sprintf("Load LBS configure file return error: %v", err))
	}

	GlsHost = confObj.MustValue(_GLS_SECTION, "lbs_host", _DEFAULT_GLS_HOST)
	InQueueSize = confObj.MustInt(_GLS_SECTION, "inqueue_size", _DEFAULT_IN_SIZE)
	OutQueueSize = confObj.MustInt(_GLS_SECTION, "outqueue_size", _DEFAULT_OUT_SIZE)
}
