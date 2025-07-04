package gologram

import (
	"bytes"
	"errors"
	"net/http"
	"runtime"
	"strconv"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type Err struct {
	err    error
	stack  []string
	code   int
	msg    string
	Type   string
	Values map[string]string
}

func NewBadRequest(err error) *Err {
	return newErr(err).WithCode(http.StatusBadRequest)
}

func NewUnauthorized(err error) *Err {
	return newErr(err).WithCode(http.StatusUnauthorized)
}

func NewForbidden(err error) *Err {
	if err == nil {
		err = errors.New("forbidden")
	}
	return newErr(err).WithCode(http.StatusForbidden)
}

func NewTypeErr(msg, t string) *Err {
	return newErr(errors.New(msg)).WithType(t)
}

func NewErr(err error) *Err {
	return newErr(err)
}

func NewTextErr(s string) *Err {
	return newErr(errors.New(s))
}

func newErr(err error) *Err {
	if err == nil {
		return nil
	}
	var ex *Err
	if errors.As(err, &ex) {
		ex.stack = append(ex.stack, getNestedPath())
		return ex
	}
	return &Err{
		stack: []string{getPath(err)},
		err:   err,
		msg:   err.Error(),
	}
}

func (e *Err) WithCode(code int) *Err {
	e.code = code
	return e
}

func (e *Err) WithStack(stack []string) *Err {
	e.stack = stack
	return e
}

func (e *Err) WithType(t string) *Err {
	e.Type = t
	return e
}

func (e *Err) WithValue(k, v string) *Err {
	if e.Values == nil {
		e.Values = make(map[string]string)
	}
	e.Values[k] = v
	return e
}

func (e *Err) Code() int {
	return e.code
}

func (e *Err) Msg() string {
	return e.msg
}

func (e *Err) Error() string {
	return e.msg
}

func (e *Err) String() string {
	if e.stack != nil && len(e.stack) > 0 {
		var sb *bytes.Buffer
		for _, s := range e.stack {
			if sb == nil {
				sb = bytes.NewBuffer(ToBytes(s))
			} else {
				sb.WriteString(s)
			}
		}
		return sb.String()
	}
	return ""
}

func getPath(err error) string {
	pc, filename, line, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc)
	var fnStr string
	if fn != nil {
		fnStr = fn.Name()
	}
	return "[error] in " + fnStr + "[" + filename + ":" + strconv.Itoa(line) + "] " + err.Error()
}

func getNestedPath() string {
	pc, filename, line, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc)
	var fnStr string
	if fn != nil {
		fnStr = fn.Name()
	}
	return "\n\t" + fnStr + "[" + filename + ":" + strconv.Itoa(line) + "]"
}
