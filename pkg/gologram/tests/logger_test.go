package tests

import (
	"errors"
	"gologram"
	"testing"
	"time"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

func TestLogger(t *testing.T) {
	debug := gologram.NewConfig().
		WithLevel(gologram.DEBUG).
		WithFormat(gologram.JSON).
		WithTimeFormat(gologram.ISO8601).Build().Named("Debug")
	info := gologram.NewConfig().
		WithLevel(gologram.INFO).
		WithFormat(gologram.JSON).
		WithTimeFormat(gologram.ISO8601).Build().Named("Info")
	warn := gologram.NewConfig().
		WithLevel(gologram.WARN).
		WithFormat(gologram.JSON).
		WithTimeFormat(gologram.ISO8601).Build().Named("Warn")
	err := gologram.NewConfig().
		WithLevel(gologram.ERROR).
		WithFormat(gologram.JSON).
		WithTimeFormat(gologram.ISO8601).Build().Named("Error")
	for i := 0; i < 100; i++ {
		go debug.Debug("Hello, World! DEBUG")
		go info.Info("Hello, World! INFO")
		go warn.Warn("Hello, World! WARN")
		go err.Error("Hello, World! ERROR", gologram.NewErr(errors.New("ereer")))
		time.Sleep(5 * time.Second)
	}
}
