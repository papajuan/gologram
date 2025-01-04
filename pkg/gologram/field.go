package gologram

import (
	"fmt"
	"strconv"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type Field struct {
	name string
	val  string
}

func NewErrField(err error) *Field {
	return &Field{name: "error", val: err.Error()}
}

func NewField(name string) *Field {
	return &Field{name: name}
}

func (f *Field) WithStringer(val fmt.Stringer) *Field {
	f.val = val.String()
	return f
}

func (f *Field) WithString(val string) *Field {
	f.val = val
	return f
}

func (f *Field) WithInt(val int) *Field {
	f.val = strconv.Itoa(val)
	return f
}

func (f *Field) WithInt64(val int64) *Field {
	f.val = strconv.FormatInt(val, 10)
	return f
}

func (f *Field) WithUint(val uint) *Field {
	f.val = strconv.FormatUint(uint64(val), 10)
	return f
}

func (f *Field) WithUint64(val uint64) *Field {
	f.val = strconv.FormatUint(val, 10)
	return f
}

// WithFloat64 See https://golang.org/pkg/strconv/#FormatFloat for fmt options
func (f *Field) WithFloat64(val float64, fmt byte) *Field {
	f.val = strconv.FormatFloat(val, fmt, -1, 32)
	return f
}

func (f *Field) WithBool(val bool) *Field {
	f.val = strconv.FormatBool(val)
	return f
}

func (f *Field) String() string {
	return `"` + f.name + `":` + `"` + f.val + `"`
}
