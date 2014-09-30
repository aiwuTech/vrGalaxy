// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package conf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego/logs"
)

const (
	_GLS_SECTION              = "gls"
	_LOG_SECTION              = "log"
	_DEFAULT_GLS_HOST         = ":8888"
	_DEFAULT_IN_SIZE          = 64
	_DEFAULT_OUT_SIZE         = 15
	_DEFAULT_LOG_NAME         = "gls.log"
	_DEFAULT_LOG_CACHE        = 10000
	_DEFAULT_LOG_LINE         = 100000
	_DEFAULT_LOG_DAILY_ROTATE = true
	_DEFAULT_LOG_DAYS         = 7
	_DEFAULT_LOG_ROTATE       = true
	_DEFAULT_LOG_LEVEL        = logs.LevelDebug
	_DEFAULT_LOG_CONSOLE      = true
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

	//log
	LogName     string
	MaxCache    int64
	MaxLines    int64
	MaxSize     int64
	DailyRotate bool
	MaxDays     int64
	LogRotate   bool
	LogLevel    int
	LogConsole  bool
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

	// log
	LogName = confObj.MustValue(_LOG_SECTION, "log_name", _DEFAULT_LOG_NAME)
	MaxCache = confObj.MustInt64(_LOG_SECTION, "max_cache", _DEFAULT_LOG_CACHE)
	MaxLines = confObj.MustInt64(_LOG_SECTION, "max_line", _DEFAULT_LOG_LINE)
	DailyRotate = confObj.MustBool(_LOG_SECTION, "daily_rotate", _DEFAULT_LOG_DAILY_ROTATE)
	MaxDays = confObj.MustInt64(_LOG_SECTION, "max_days", _DEFAULT_LOG_DAYS)
	LogRotate = confObj.MustBool(_LOG_SECTION, "log_rotate", _DEFAULT_LOG_ROTATE)
	LogLevel = confObj.MustInt(_LOG_SECTION, "log_level", _DEFAULT_LOG_LEVEL)
	LogConsole = confObj.MustBool(_LOG_SECTION, "log_console", _DEFAULT_LOG_CONSOLE)
}
