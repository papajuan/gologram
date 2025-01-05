package gologram

import (
	"encoding/json"
	"strings"
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
	Fields     []*Field `json:"fields"`
}

func NewOutput(ts, caller, severity, msg string, err *Err, fields ...*Field) *Output {
	result := &Output{
		Timestamp: ts,
		Severity:  severity,
		Message:   msg,
		Caller:    caller,
		Fields:    fields,
	}
	if err != nil {
		result.Stacktrace = err.String()
	}
	return result
}

func (o *Output) JsonString() []byte {
	res, err := json.Marshal(o)
	if err != nil {
		return []byte("{}\n")
	}
	res = append(res, '\n')
	return res
}

func (o *Output) ConsoleString() []byte {
	var builder strings.Builder
	builder.WriteString(o.Timestamp)
	builder.WriteString("\t")
	builder.WriteString(o.Severity)
	if o.Caller != "" {
		builder.WriteString("\t")
		builder.WriteString(o.Caller)
	}
	builder.WriteString("\t")
	builder.WriteString(o.Message)
	if o.Stacktrace != "" {
		builder.WriteString("\n")
		builder.WriteString(o.Stacktrace)
	}
	if o.Fields != nil && len(o.Fields) > 0 {
		builder.WriteString("\t{")
		for i, field := range o.Fields {
			builder.WriteString(field.String())
			if i < len(o.Fields)-1 {
				builder.WriteString(",")
			}
		}
		builder.WriteRune('}')
	}
	builder.WriteRune('\n')
	return []byte(builder.String())
}
