package tests

import (
	"errors"
	"fmt"
	"gologram"
	"net/http"
	"strings"
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

func BenchmarkConcat(b *testing.B) {
	b.Run("Concat with +10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res := "Hello, " + "World!"
			res += "1"
			res += "2"
			res += "3"
			res += "4"
			res += "5"
			res += "6"
			res += "7"
			res += "8"
			res += "9"
			res += "10"
		}
	})
	w := "World!"
	one := "1"
	two := "2"
	three := "3"
	four := "4"
	five := "5"
	six := "6"
	seven := "7"
	eight := "8"
	nine := "9"
	ten := "10"
	b.Run("Concat with + 10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = "Hello, " + w + one + two + three + four + five + six + seven + eight + nine + ten
		}
	})
	b.Run("Concat with fmt 10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("Hello, %s%s%s%s%s%s%s%s%s%s%s", w, one, two, three, four, five, six, seven, eight, nine, ten)
		}
	})
	b.Run("Concat with strings.Builder 10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var builder strings.Builder
			builder.WriteString("Hello, ")
			builder.WriteString("World!")
			builder.WriteString("1")
			builder.WriteString("2")
			builder.WriteString("3")
			builder.WriteString("4")
			builder.WriteString("5")
			builder.WriteString("6")
			builder.WriteString("7")
			builder.WriteString("8")
			builder.WriteString("9")
			builder.WriteString("10")
		}
	})
}
