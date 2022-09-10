package structformat

type StringFormatter string

var _ Formatter = (*StringFormatter)(nil)

// Return a StringFormatter for a string. New lines will be indented according to
// the current indent level of the formatter.
func String(s string) StringFormatter {
	return StringFormatter(s)
}

func (f StringFormatter) StructFormat(w Writer) (n int, err error) {
	return w.Write([]byte(f))
}
