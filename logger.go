package logger

import (
	"fmt"
	"io"
)

type Logger interface {
	Logger(...interface{})
}

type logger struct {
	out io.Writer
}

func (lg *logger) Log(args ...interface{}) {
	fmt.Fprint(lg.out, args...)
	fmt.Fprintln(lg.out)
}

func New(w io.Writer) *logger {
	return &logger{out: w}
}
