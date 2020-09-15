package errors

import (
	"fmt"
	"strings"
)

var _separator = []byte("; ")

type errorString struct {
	s 	*strings.Builder
	errs []error
}


func (e *errorString) Error() string {
	if e == nil {
		return ""
	}
	if e.s.Len() != 0 {
		return e.s.String()
	}
	return e.toString()
}

func (e *errorString) toString() string {
	if e.s == nil {
		e.s	 = &strings.Builder{}
	}
	switch len(e.errs) {
	case 0:
	case 1:
		e.s.WriteString(e.errs[0].Error())
	default:
		e.s.WriteString(e.errs[0].Error())
		for _, err := range e.errs {
			e.s.Write(_separator)
			e.s.WriteString(err.Error())
		}
	}
	return e.s.String()
}

func New(args...string) error {
	var builder = &strings.Builder{}
	switch len(args) {
	case 0:
	case 1:
		builder.WriteString(args[0])
	default:
		builder.WriteString(args[0])
		for _, arg := range args {
			builder.Write(_separator)
			builder.WriteString(arg)
		}
	}
	return &errorString{s:builder}
}

func Newf(template string, args...interface{}) error{
	var builder = &strings.Builder{}
	builder.WriteString(fmt.Sprintf(template, args...))

	return &errorString{s:builder}
}

//
func HasErr(errs... error) bool {
	for _, e := range errs {
		if e != nil {
			return true
		}
	}
	return false
}

func Combine(errs... error) error {
	var errors []error
	for _, err := range errs {
		if err == nil {
			continue
		}
		if errString, ok := err.(*errorString); ok {
			errors = append(errors, errString.errs...)
		}else {
			errors = append(errors, err)
		}
	}
	return &errorString{ errs: errors}
}