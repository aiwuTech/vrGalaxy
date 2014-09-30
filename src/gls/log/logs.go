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
	glsLog *logs.BeeLogger
	consoleLog *logs.BeeLogger
)

func init() {

	glsLog = logs.NewLogger(conf.MaxCache)

	logConf := fmt.Sprintf("{\"filename\":\"%s\",\"maxlines\":%v,\"maxsize\":%v,\"daily\":%v,\"maxdays\":%v,\"rotate\":%v,\"level\":%v}",
		conf.LogName, conf.MaxLines, conf.MaxSize, conf.DailyRotate, conf.MaxDays, conf.LogRotate, conf.LogLevel)
	err := glsLog.SetLogger("file", logConf)
	if err != nil {
		panic(fmt.Sprintf("start file log return error: %v", err))
	}
	glsLog.EnableFuncCallDepth(true)
	glsLog.SetLogFuncCallDepth(4)

	consoleLog = logs.NewLogger(conf.MaxCache)
	err = consoleLog.SetLogger("console", `{"level":7}`)
	if err != nil {
		panic(fmt.Sprintf("start console log return error: %v", err))
	}
	consoleLog.EnableFuncCallDepth(true)
	consoleLog.SetLogFuncCallDepth(4)
}

// trace, debug
func T(format string, v ...interface {}) {
	glsLog.Trace(format, v...)
	if conf.LogConsole {
		consoleLog.Trace(format, v...)
	}
}

// info
func I(format string, v ...interface {}) {
	glsLog.Info(format, v...)
	if conf.LogConsole {
		consoleLog.Info(format, v...)
	}
}

// warning
func W(format string, v ...interface {}) {
	glsLog.Warn(format, v...)
	if conf.LogConsole {
		consoleLog.Warn(format, v...)
	}
}

// error
func E(format string, v ...interface {}) {
	glsLog.Error(format, v...)
	if conf.LogConsole {
		consoleLog.Error(format, v...)
	}
}

// fatal
func F(format string, v ...interface {}) {
	glsLog.Critical(format, v...)
	if conf.LogConsole {
		consoleLog.Critical(format, v...)
	}
}
