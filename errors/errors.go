package errors

import (
	"fmt"
	"strings"
)

type errorString struct {
	s 	string
}


func (e *errorString) Error() string {
	return e.s
}

func New(args...string) error {
	return &errorString{s:strings.Join(args, ",")}
}

func Newf(template string, args...interface{}) error{
	return &errorString{s:fmt.Sprintf(template, args...)}
}