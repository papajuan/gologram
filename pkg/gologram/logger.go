package gologram

import (
	"gologram/buffer"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type Logger struct {
	name           string
	level          *logLevel
	timeFormatFunc TimeFormatFunc
	format         Format
}

func newLogger(config *Config) *Logger {
	buffer.Initialize()
	return &Logger{
		level:          config.level,
		timeFormatFunc: config.timeFormatFunc,
		format:         config.format,
	}
}

func (l *Logger) Named(msg string) *Logger {
	l.name = msg
	return l
}

func (l *Logger) Trace(msg string, fields ...*Field) {
	buffer.Stdout().Write(l.output(TRACE, msg, nil, fields...))
}

func (l *Logger) Debug(msg string, fields ...*Field) {
	if l.level.level <= DEBUG {
		buffer.Stdout().Write(l.output(DEBUG, msg, nil, fields...))
	}
}

func (l *Logger) Info(msg string, fields ...*Field) {
	if l.level.level <= INFO {
		buffer.Stdout().Write(l.output(INFO, msg, nil, fields...))
	}
}

func (l *Logger) Warn(msg string, fields ...*Field) {
	if l.level.level <= WARN {
		buffer.Stderr().Write(l.output(WARN, msg, nil, fields...))
	}
}

func (l *Logger) Error(msg string, err *Err, fields ...*Field) {
	if l.level.level <= ERROR {
		buffer.Stderr().Write(l.output(ERROR, msg, err, fields...))
	}
}

func (l *Logger) output(lev Level, msg string, err *Err, fields ...*Field) []byte {
	switch l.format {
	case JSON:
		return NewOutput(l.timeFormatFunc(), l.name, lev.String(), msg, err, fields...).JsonString()
	default:
		return NewOutput(l.timeFormatFunc(), l.name, lev.StringColored(), msg, err, fields...).ConsoleString()
	}
}
