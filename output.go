package gologram

import (
	"bytes"
	"strings"
	"unsafe"
)

/**
 * @author  papajuan
 * @date    1/5/2025
 **/

type Output struct {
	Timestamp  string   `json:"timestamp"`
	Severity   string   `json:"severity"`
	Message    string   `json:"message"`
	Stacktrace string   `json:"stacktrace"`
	Caller     string   `json:"caller"`
	callerTab  string   `json:"-"`
	Fields     []*Field `json:"fields"`
}

func newOutput(ts, caller, severity, msg string, err *Err, fields ...*Field) *Output {
	var callerTab string
	if caller == "" {
		callerTab = "    "
	} else if len(caller) < 15 {
		l := 15 - len(caller)
		sb := make([]byte, 0, l)
		for i := 0; i < l; i++ {
			sb = append(sb, ' ')
		}
		callerTab = unsafe.String(unsafe.SliceData(sb), l)
	} else {
		callerTab = " "
	}
	result := &Output{
		Timestamp: ts,
		Severity:  severity,
		Message:   strings.ReplaceAll(msg, `"`, `'`),
		Caller:    caller,
		Fields:    fields,
		callerTab: callerTab,
	}
	if err != nil {
		result.Stacktrace = strings.ReplaceAll(err.String(), `"`, `'`)
	}
	return result
}

func (o *Output) jsonString() []byte {
	return ToBytes(`{"timestamp":"` + o.Timestamp +
		`","severity":"` + o.Severity +
		`","stacktrace":"` + o.Stacktrace +
		`","caller":"` + o.Caller +
		`","message":"` + o.Caller + " " + o.Message + " " + o.StringFields() + "\"}\n")
}

func (o *Output) StringFields() string {
	var result string
	if o.Fields != nil {
		if l := len(o.Fields); l > 0 {
			var sb bytes.Buffer
			for i, field := range o.Fields {
				sb.Write(ToBytes(field.String()))
				if i < l-1 {
					sb.WriteByte(',')
				}
			}
			result = sb.String()
		}
	}
	return result
}

func (o *Output) consoleString() []byte {
	sb := bytes.NewBuffer(ToBytes(o.Timestamp + "\t" + o.Severity))
	if o.Caller != "" {
		sb.Write(ToBytes("\t" + o.Caller))
	}
	sb.Write(ToBytes(o.callerTab + o.Message))
	if o.Stacktrace != "" {
		sb.Write(ToBytes("\n" + o.Stacktrace))
	}
	sb.WriteString(" " + o.StringFields())
	sb.WriteByte('\n')
	return sb.Bytes()
}

func ToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
