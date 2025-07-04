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
	RFC3339Nano  = "RFC3339Nano"
	ISO8601      = "ISO8601"
	ANSIC        = "ANSIC"
	Nanoseconds  = "nanos"
	Microseconds = "micros"
	Milliseconds = "millis"
	Seconds      = "secs"
)

type TimeFormatFunc func() string

type config struct {
	level          *logLevel
	format         Format
	timeFormatFunc TimeFormatFunc
}

func NewConfig() *config {
	return &config{}
}

func (c *config) WithLevel(l Level) *config {
	c.level = newLogLevel(l)
	return c
}

func (c *config) WithFormat(f Format) *config {
	c.format = f
	return c
}

func (c *config) WithTimeFormat(f string) *config {
	c.timeFormatFunc = getTimeFormatFunc(f)
	return c
}

func (c *config) Build() *logger {
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
		case ISO8601:
			return func() string { return time.Now().Format("2006-01-02T15:04:05.000Z0700") }
		case RFC3339Nano:
			return func() string { return time.Now().Format(time.RFC3339Nano) }
		case ANSIC:
			return func() string { return time.Now().Format("2006-01-02 15:04:05 MST") }
		default:
			return func() string { return time.Now().Format(timeFormat) }
		}
	}
}
