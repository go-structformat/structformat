package structformat

import "fmt"

type PrintfFormatter struct {
	Fmt  string
	Args []interface{}
}

var _ Formatter = (*PrintfFormatter)(nil)

// Return a PrintfFormatter that apply Printf template to the arguments and format them as a string.
func Printf(format string, args ...interface{}) *PrintfFormatter {
	return &PrintfFormatter{Fmt: format, Args: args}
}

func (f *PrintfFormatter) StructFormat(w Writer) (n int, err error) {
	return fmt.Fprintf(w, f.Fmt, f.Args...)
}
