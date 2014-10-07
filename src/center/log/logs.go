// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package log

import (
	"github.com/astaxie/beego/logs"
	"gls/conf"
	"fmt"
)

var (
	_log *logs.BeeLogger
	_clog *logs.BeeLogger
)

func init() {

	_log = logs.NewLogger(conf.MaxCache)

	logConf := fmt.Sprintf("{\"filename\":\"%s\",\"maxlines\":%v,\"maxsize\":%v,\"daily\":%v,\"maxdays\":%v,\"rotate\":%v,\"level\":%v}",
		conf.LogName, conf.MaxLines, conf.MaxSize, conf.DailyRotate, conf.MaxDays, conf.LogRotate, conf.LogLevel)
	err := _log.SetLogger("file", logConf)
	if err != nil {
		panic(fmt.Sprintf("start file log return error: %v", err))
	}
	_log.EnableFuncCallDepth(true)
	_log.SetLogFuncCallDepth(4)

	_clog = logs.NewLogger(conf.MaxCache)
	err = _clog.SetLogger("console", `{"level":7}`)
	if err != nil {
		panic(fmt.Sprintf("start console log return error: %v", err))
	}
	_clog.EnableFuncCallDepth(true)
	_clog.SetLogFuncCallDepth(4)
}

// trace, debug
func T(format string, v ...interface {}) {
	_log.Trace(format, v...)
	if conf.LogConsole {
		_clog.Trace(format, v...)
	}
}

// info
func I(format string, v ...interface {}) {
	_log.Info(format, v...)
	if conf.LogConsole {
		_clog.Info(format, v...)
	}
}

// warning
func W(format string, v ...interface {}) {
	_log.Warn(format, v...)
	if conf.LogConsole {
		_clog.Warn(format, v...)
	}
}

// error
func E(format string, v ...interface {}) {
	_log.Error(format, v...)
	if conf.LogConsole {
		_clog.Error(format, v...)
	}
}

// fatal
func F(format string, v ...interface {}) {
	_log.Critical(format, v...)
	if conf.LogConsole {
		_clog.Critical(format, v...)
	}
}
