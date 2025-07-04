package gologram

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type logLevel struct {
	level       Level
	nameColored string
	name        string
}

func NewLogLevel(level string) Level {
	switch level {
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "TRACE":
		return TRACE
	default:
		return DEBUG
	}
}

func newLogLevel(l Level) *logLevel {
	return &logLevel{level: l, nameColored: l.stringColored(), name: l.String()}
}

type Level uint8

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)

func (l Level) getSeverity() string {
	switch l {
	case TRACE, DEBUG:
		return "DEBUG"
	case WARN:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "INFO"
	}
}

// ANSI color codes
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	grey   = "\033[90m"

	debug = blue + "DEBUG" + reset
	info  = green + "INFO" + reset
	warn  = yellow + "WARN" + reset
	er    = red + "ERROR" + reset
	trace = grey + "TRACE" + reset
)

func (l Level) stringColored() string {
	switch l {
	case DEBUG:
		return debug
	case INFO:
		return info
	case WARN:
		return warn
	case ERROR:
		return er
	case TRACE:
		return trace
	default:
		return ""
	}
}

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case TRACE:
		return "TRACE"
	default:
		return ""
	}
}
