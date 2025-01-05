package tests

import (
	"errors"
	"gologram"
	"net/http"
	"testing"
	"time"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

var l = gologram.NewConfig().
	WithLevel(gologram.TRACE).
	WithFormat(gologram.CONSOLE).
	WithTimeFormat(gologram.ISO8601).Build().Named("BenchTest")

func BenchmarkLoggerError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l.Error("Hello, World! ERROR", gologram.NewErr(errors.New("ereer")).WithCode(http.StatusUnprocessableEntity))
	}
}

func BenchmarkLoggerInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l.Info("Hello, World! ERROR", gologram.NewField("testname").WithString("testvalue"))
	}
}

func TestLogger(t *testing.T) {
	debug := gologram.NewConfig().
		WithLevel(gologram.DEBUG).
		WithFormat(gologram.CONSOLE).
		WithTimeFormat(gologram.ISO8601).Build().Named("Debug")
	info := gologram.NewConfig().
		WithLevel(gologram.INFO).
		WithFormat(gologram.CONSOLE).
		WithTimeFormat(gologram.ISO8601).Build().Named("Info")
	warn := gologram.NewConfig().
		WithLevel(gologram.WARN).
		WithFormat(gologram.CONSOLE).
		WithTimeFormat(gologram.ISO8601).Build().Named("Warn")
	err := gologram.NewConfig().
		WithLevel(gologram.ERROR).
		WithFormat(gologram.CONSOLE).
		WithTimeFormat(gologram.ISO8601).Build().Named("Error")
	for i := 0; i < 100; i++ {
		go debug.Debug("Hello, World! DEBUG")
		go info.Info("Hello, World! INFO")
		go warn.Warn("Hello, World! WARN")
		go err.Error("Hello, World! ERROR", gologram.NewErr(errors.New("ereer")).WithCode(http.StatusUnprocessableEntity))
		time.Sleep(5 * time.Second)
	}
}
