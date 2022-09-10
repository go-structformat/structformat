package structformat

import "fmt"

type PrintfFormatter struct {
	Fmt  string
	Args []interface{}
}

var _ Formatter = (*PrintfFormatter)(nil)

func Printf(format string, args ...interface{}) *PrintfFormatter {
	return &PrintfFormatter{Fmt: format, Args: args}
}

func (f *PrintfFormatter) StructFormat() NestedLines {
	return NestedLines{fmt.Sprintf(f.Fmt, f.Args...)}
}
