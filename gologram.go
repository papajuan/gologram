package gologram

import (
	"fmt"
	"os"
	"runtime"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

var (
	cfg *config
	log *logger
)

func init() {
	ll := NewLogLevel(os.Getenv("LOG_LEVEL"))
	cfg = NewConfig().
		WithLevel(ll).
		WithFormat(NewLogFormat(os.Getenv("LOG_FORMAT"))).
		WithTimeFormat(os.Getenv("LOG_TIME_FORMAT"))
	log = New("Panic")
}

func Sync() {
	err := stdout().Sync()
	if err != nil {
		println(err)
	}
	err = stderr().Sync()
	if err != nil {
		println(err)
	}
}

func New(name string) *logger {
	return cfg.Build().Named(name)
}

func Safe(f func()) {
	defer func() {
		if r := recover(); r != nil {
			// TODO improve
			stackBuf := make([]byte, 1024)
			length := runtime.Stack(stackBuf, false)
			stack := make([]string, length)
			for i, barr := range stackBuf[:length] {
				stack[i] = string(barr)
			}
			log.Error("Recovered from panic", NewErr(fmt.Errorf("%v", r)).WithStack(stack))
		}
	}()
	f()
}
