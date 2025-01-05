package gologram

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type Err struct {
	err   error
	stack []string
	code  int
	msg   string
}

func NewBadRequest(err error) *Err {
	return newErr(err).WithCode(http.StatusBadRequest)
}

func NewErr(err error) *Err {
	return newErr(err)
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

func (e *Err) Code() int {
	return e.code
}

func (e *Err) Error() string {
	return e.msg
}

func (e *Err) String() string {
	var builder strings.Builder
	for _, s := range e.stack {
		builder.WriteString(s)
	}
	return builder.String()
}

func getPath(err error) string {
	pc, filename, line, _ := runtime.Caller(3)
	return fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
}

func getNestedPath() string {
	pc, filename, line, _ := runtime.Caller(3)
	return fmt.Sprintf("\n\t%s[%s:%d]", runtime.FuncForPC(pc).Name(), filename, line)
}
