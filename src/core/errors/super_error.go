package errors

import (
	"runtime"
)

type superError struct {
	file string
	line int
}

func newSuperError() *superError {
	_, file, line, _ := runtime.Caller(2)
	return &superError{file: file, line: line}
}

func (instance superError) File() string {
	return instance.file
}

func (instance superError) Line() int {
	return instance.line
}

func (instance superError) LogLevel() int {
	return InfoLevel
}
