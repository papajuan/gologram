package gologram

import (
	"gologram/buffer"
	"strings"
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
	buffer.Stdout().Write(l.output(msg, nil, fields...))
}

func (l *Logger) Debug(msg string, fields ...*Field) {
	if l.level.level >= DEBUG {
		buffer.Stdout().Write(l.output(msg, nil, fields...))
	}
}

func (l *Logger) Info(msg string, fields ...*Field) {
	if l.level.level >= INFO {
		buffer.Stdout().Write(l.output(msg, nil, fields...))
	}
}

func (l *Logger) Warn(msg string, fields ...*Field) {
	if l.level.level >= WARN {
		buffer.Stderr().Write(l.output(msg, nil, fields...))
	}
}

func (l *Logger) Error(msg string, err *Err, fields ...*Field) {
	if l.level.level >= ERROR {
		buffer.Stderr().Write(l.output(msg, err, fields...))
	}
}

func (l *Logger) output(msg string, err *Err, fields ...*Field) []byte {
	var builder strings.Builder
	switch l.format {
	case CONSOLE:
		builder.WriteString(l.timeFormatFunc())
		builder.WriteString("\t")
		builder.WriteString(l.level.nameColored)
		if l.name != "" {
			builder.WriteString("\t")
			builder.WriteString(l.name)
		}
		builder.WriteString("\t")
		if err != nil {
			// TODO implement error writing
		}
		builder.WriteString(msg)
		if fields != nil && len(fields) > 0 {
			builder.WriteString("\t{")
			for i, field := range fields {
				builder.WriteString(field.String())
				if i < len(fields)-1 {
					builder.WriteString(",")
				}
			}
			builder.WriteRune('}')
		}
	case JSON:
		builder.WriteRune('{')
		builder.WriteString(`"timestamp":"` + l.timeFormatFunc() + `",`)
		builder.WriteString(`"severity":"` + l.level.name + `",`)
		builder.Write([]byte(`"message":"` + msg + `",`))
		if err != nil {
			// TODO implement error and stack trace writing
		}
		builder.WriteRune('}')
	}
	builder.WriteString("\n")
	return []byte(builder.String())
}
