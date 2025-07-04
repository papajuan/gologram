package gologram

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

var (
	replacer = strings.NewReplacer(`\n`, "", `"`, `'`, `\`, "")
)

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

func (f *Field) WithStringArr(val []string) *Field {
	if val == nil {
		f.val = "null"
	} else if len(val) == 0 {
		f.val = "[]"
	} else {
		result := bytes.NewBuffer([]byte{'['})
		for _, s := range val {
			result.Write(ToBytes(`"` + s + `",`))
		}
		f.val = result.String()[:result.Len()-1] + "]"
	}
	return f
}

func (f *Field) WithByteArr(val []byte) *Field {
	if val == nil || len(val) == 0 {
		f.val = "null"
	} else {
		f.val = fmt.Sprintf(`%s`, val)
	}
	return f
}

func (f *Field) WithAny(val interface{}) *Field {
	f.val = fmt.Sprintf("%v", val)
	return f
}

func (f *Field) WithByte(b byte) *Field {
	f.val = strconv.FormatUint(uint64(b), 10)
	return f
}

func (f *Field) String() string {

	return `'` + f.name + `':'` + replacer.Replace(f.val) + `'`
}

func StringField(name, val string) *Field {
	return &Field{name: name, val: val}
}

func ByteStringField(name string, val []byte) *Field {
	return &Field{name: name, val: string(val)}
}

func AnyField(name string, val interface{}) *Field {
	return &Field{name: name, val: fmt.Sprintf("%+v", val)}
}
