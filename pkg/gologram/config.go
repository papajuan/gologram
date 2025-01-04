package gologram

import (
	"strconv"
	"time"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

const (
	ISO8601      = "2006-01-02T15:04:05.000Z0700"
	Nanoseconds  = "nanos"
	Microseconds = "micros"
	Milliseconds = "millis"
	Seconds      = "secs"
)

type TimeFormatFunc func() string

type Config struct {
	level          *logLevel
	format         Format
	timeFormatFunc TimeFormatFunc
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithLevel(l Level) *Config {
	c.level = newLogLevel(l)
	return c
}

func (c *Config) WithFormat(f Format) *Config {
	c.format = f
	return c
}

func (c *Config) WithTimeFormat(f string) *Config {
	c.timeFormatFunc = getTimeFormatFunc(f)
	return c
}

func (c *Config) Build() *Logger {
	if c.level == nil {
		c.level = newLogLevel(DEBUG)
	}
	if c.timeFormatFunc == nil {
		c.timeFormatFunc = getTimeFormatFunc("")
	}
	return newLogger(c)
}

func getTimeFormatFunc(timeFormat string) TimeFormatFunc {
	if timeFormat == "" {
		return func() string { return time.Now().String() }
	} else {
		switch timeFormat {
		case Nanoseconds:
			return func() string { return strconv.FormatInt(time.Now().UnixNano(), 10) }
		case Microseconds:
			return func() string { return strconv.FormatInt(time.Now().UnixMicro(), 10) }
		case Milliseconds:
			return func() string { return strconv.FormatInt(time.Now().UnixMilli(), 10) }
		case Seconds:
			return func() string { return strconv.FormatInt(time.Now().Unix(), 10) }
		default:
			return func() string { return time.Now().Format(timeFormat) }
		}
	}
}
