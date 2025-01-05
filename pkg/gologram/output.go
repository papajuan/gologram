package gologram

import (
	"strings"
)

/**
 * @author  papajuan
 * @date    1/5/2025
 **/

type output struct {
	timestamp  string
	severity   string
	caller     string
	message    string
	stacktrace string
	fields     []*Field
}

func NewOutput(format Format, ts, caller, severity, msg string, err *Err, fields ...*Field) *output {
	var res output
	if format == JSON {
		res = output{fields: []*Field{
			NewField("timestamp").WithString(ts),
			NewField("severity").WithString(severity),
			NewField("caller").WithString(caller),
			NewField("message").WithString(msg),
		}}
		if err != nil {
			res.fields = append(res.fields, NewField("stacktrace").WithString(err.String()))
		}
		res.fields = append(res.fields, fields...)
	} else {
		res = output{
			timestamp: ts,
			severity:  severity,
			caller:    caller,
			message:   msg,
			fields:    fields,
		}
		if err != nil {
			res.stacktrace = err.String()
		}
	}
	return &res
}

func (o *output) JsonString() []byte {
	var builder strings.Builder
	if o.fields != nil && len(o.fields) > 0 {
		for i, field := range o.fields {
			builder.WriteString(field.String())
			if i < len(o.fields)-1 {
				builder.WriteString(",")
			}
		}
	}
	return []byte("{" + builder.String() + "}\n")
}

func (o *output) ConsoleString() []byte {
	var builder strings.Builder
	builder.WriteString(o.timestamp + "\t" + o.severity)
	if o.caller != "" {
		builder.WriteString("\t" + o.caller)
	}
	builder.WriteString("\t" + o.message)
	if o.stacktrace != "" {
		builder.WriteString("\n" + o.stacktrace)
	}
	if o.fields != nil && len(o.fields) > 0 {
		builder.WriteString("\t")
		for i, field := range o.fields {
			builder.WriteString(field.String())
			if i < len(o.fields)-1 {
				builder.WriteString(",")
			}
		}
	}
	builder.WriteString("\n")
	return []byte(builder.String())
}
