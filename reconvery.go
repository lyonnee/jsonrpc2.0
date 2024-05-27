package jsonrpc

import (
	"io"
	"log"
	"os"
	"runtime/debug"
	"time"
)

var DefaultErrorWriter io.Writer = os.Stderr

// 定义一个全局变量来存储用户自定义的异常处理函数
var recoverHandler = defaultRecoverHandler

// 定义默认的异常处理函数
var defaultRecoverHandler = func(err any) {
	var logger *log.Logger
	if DefaultErrorWriter != nil {
		logger = log.New(DefaultErrorWriter, "\n\n\x1b[31m", log.LstdFlags)
	}

	if logger != nil {
		stack := debug.Stack()
		logger.Printf("[Recovery] %s panic recovered:\n%s\n%s",
			timeFormat(time.Now()), err, string(stack))
	}
}

// SetRecoverHandler 设置自定义的异常处理函数
func SetRecoverHandler(handler func(err any)) {
	if handler != nil {
		recoverHandler = handler
	}
}

// timeFormat returns a customized time string for logger.
func timeFormat(t time.Time) string {
	return t.Format("2006/01/02 - 15:04:05")
}
