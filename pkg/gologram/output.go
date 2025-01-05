package gologram

import (
	"strings"
)

/**
 * @author  papajuan
 * @date    1/5/2025
 **/

type output struct {
	fields []*Field
}

func NewOutput(ts, caller, severity, msg string, err *Err, fields ...*Field) *output {
	res := output{[]*Field{
		NewField("timestamp").WithString(ts),
		NewField("severity").WithString(severity),
		NewField("caller").WithString(caller),
		NewField("message").WithString(msg),
	}}
	if err != nil {
		res.fields = append(res.fields, NewField("stacktrace").WithString(err.String()))
	}
	res.fields = append(res.fields, fields...)
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
	builder.WriteString(o.fields[0].val + "\t" + o.fields[1].val)
	if o.fields[2].val != "" {
		builder.WriteString("\t" + o.fields[2].val)
	}
	builder.WriteString("\t" + o.fields[3].val)
	if o.fields[4].val != "" {
		builder.WriteString("\n" + o.fields[4].val)
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
