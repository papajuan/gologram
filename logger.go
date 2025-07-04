package gologram

import (
	"fmt"
	"os"
	"rmg-market-service/gologram/buffer"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type logger struct {
	name           string
	level          *logLevel
	timeFormatFunc TimeFormatFunc
	format         Format
}

func newLogger(config *config) *logger {
	buffer.Initialize()
	return &logger{
		level:          config.level,
		timeFormatFunc: config.timeFormatFunc,
		format:         config.format,
	}
}

func (l *logger) Named(msg string) *logger {
	l.name = "[" + msg + "]"
	return l
}

func (l *logger) Trace(msg string, fields ...*Field) {
	buffer.Stdout().Write(l.output(TRACE, msg, nil, fields...))
}

func (l *logger) Debug(msg string, fields ...*Field) {
	if l.level.level <= DEBUG {
		buffer.Stdout().Write(l.output(DEBUG, msg, nil, fields...))
	}
}

func (l *logger) Info(msg string, fields ...*Field) {
	if l.level.level <= INFO {
		buffer.Stdout().Write(l.output(INFO, msg, nil, fields...))
	}
}

func (l *logger) Warn(msg string, fields ...*Field) {
	if l.level.level <= WARN {
		buffer.Stderr().Write(l.output(WARN, msg, nil, fields...))
	}
}

func (l *logger) Error(msg string, err *Err, fields ...*Field) {
	if l.level.level <= ERROR {
		buffer.Stderr().Write(l.output(ERROR, msg, err, fields...))
	}
}

func (l *logger) output(lev Level, msg string, err *Err, fields ...*Field) []byte {
	switch l.format {
	case JSON:
		return newOutput(l.timeFormatFunc(), l.name, lev.getSeverity(), msg, err, fields...).jsonString()
	default:
		return newOutput(l.timeFormatFunc(), l.name, lev.stringColored(), msg, err, fields...).consoleString()
	}
}

func (l *logger) Println(v ...interface{}) {
	buffer.Stdout().Write(l.output(INFO, fmt.Sprintf("%v\n", v...), nil))
}

func (l *logger) Printf(format string, v ...interface{}) {
	buffer.Stdout().Write(l.output(INFO, fmt.Sprintf(format, v...), nil))
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	buffer.Stderr().Write(l.output(ERROR, fmt.Sprintf(format, v...), nil))
	os.Exit(1)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	buffer.Stderr().Write(l.output(ERROR, fmt.Sprintf(format, v...), nil))
}

func (l *logger) Warnf(format string, v ...interface{}) {
	buffer.Stdout().Write(l.output(WARN, fmt.Sprintf(format, v...), nil))
}

func (l *logger) Infof(format string, v ...interface{}) {
	buffer.Stdout().Write(l.output(INFO, fmt.Sprintf(format, v...), nil))
}

func (l *logger) Debugf(format string, v ...interface{}) {
	buffer.Stdout().Write(l.output(DEBUG, fmt.Sprintf(format, v...), nil))
}
