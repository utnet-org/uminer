package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
)

type SpiderError struct {
	CodeType    CodeType
	Code        int
	Message     string
	stdError    error
	prevErr     *SpiderError // prevErr 指向上一个Err
	stack       []uintptr
	once        sync.Once
	fullMessage string
}

type stackFrame struct {
	funcName string
	file     string
	line     int
}

func (e *SpiderError) HTTPCode() int {
	if e == nil {
		return int(Unknown)
	}

	return int(e.CodeType)
}

func (e *SpiderError) HTTPSubCode() int {
	if e == nil {
		return ErrorUnknown
	}

	return e.Code
}

func (e *SpiderError) HTTPErrMsg() string {
	if e == nil {
		return codeTypeMsg[Unknown]
	}

	if msg, ok := codeTypeMsg[e.CodeType]; ok {
		return msg
	}

	return codeTypeMsg[Unknown]
}

func (e *SpiderError) Error() string {
	e.once.Do(func() {
		var buf strings.Builder
		buf.Grow(512)
		var (
			message string
			stack   []uintptr
		)
		stack = e.stack
		if e.stdError != nil {
			message = fmt.Sprintf("\n\tcodetype = %d code = %d\n\tmessage = %s\n\tstdErr = %s", e.CodeType, e.Code, e.Message, e.stdError.Error())
		} else {
			message = fmt.Sprintf("\n\tcodetype = %d code = %d\n\tmessage = %s", e.CodeType, e.Code, e.Message)
		}
		buf.WriteString(message)

		sf := stackFrame{}
		for _, v := range stack {
			funcForPc := runtime.FuncForPC(v)
			if funcForPc == nil {
				sf.file = "???"
				sf.line = 0
				sf.funcName = "???"
				buf.WriteString(fmt.Sprintf("\n\t[ %s:%d:%s]", sf.file, sf.line, sf.funcName))
				continue
			}
			sf.file, sf.line = funcForPc.FileLine(v - 1)

			//处理函数名
			sf.funcName = funcForPc.Name()

			//保证闭包函数名也能正确显示 如TestErrorf.func1:
			idx := strings.LastIndexByte(sf.funcName, '/')
			if idx != -1 {
				sf.funcName = sf.funcName[idx:]
				idx = strings.IndexByte(sf.funcName, '.')
				if idx != -1 {
					sf.funcName = strings.TrimPrefix(sf.funcName[idx:], ".")
				}
			}
			buf.WriteString(fmt.Sprintf("\n\t[ %s:%d:%s]", sf.file, sf.line, sf.funcName))
		}
		e.fullMessage = buf.String()
	})
	return e.fullMessage
}

// 如果参数error类型不为*SpiderError(error常量或自定义error类型或nil), 用于最早出错的地方, 会收集调用栈
// 如果参数error类型为*SpiderError, 不会收集调用栈
func Errorf(err error, code int) error {
	if err, ok := err.(*SpiderError); ok {
		return &SpiderError{
			CodeType: codeMsgMap[code].codeType,
			Code:     code,
			Message:  codeMsgMap[code].msg,
			prevErr:  err,
		}
	}
	newErr := new_(code, codeMsgMap[code])
	newErr.stdError = err
	return newErr
}

func Errorw(err error, codeType CodeType, code int, msg string) error {
	if err, ok := err.(*SpiderError); ok {
		return &SpiderError{
			CodeType: codeType,
			Code:     code,
			Message:  msg,
			prevErr:  err,
		}
	}
	newErr := new_(code, codeMsg{codeType: codeType, msg: msg})
	newErr.stdError = err
	return newErr
}

func new_(code int, codeMsg codeMsg) *SpiderError {
	pc := make([]uintptr, 200)
	length := runtime.Callers(3, pc)
	return &SpiderError{
		CodeType: codeMsg.codeType,
		Code:     code,
		Message:  codeMsg.msg,
		stack:    pc[:length],
	}
}

func FromError(err error) (*SpiderError, bool) {
	if se := new(SpiderError); errors.As(err, &se) {
		return se, true
	}
	return nil, false
}

func IsError(code int, err error) bool {
	if se := new(SpiderError); errors.As(err, &se) {
		return se.Code == code
	}
	return false
}
