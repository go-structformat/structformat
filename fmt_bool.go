package structformat

type BoolFormatter bool

var _ Formatter = (*BoolFormatter)(nil)

// Return a BoolFormatter that just outputs true or false.
func Bool(value bool) BoolFormatter {
	return BoolFormatter(value)
}

func (f BoolFormatter) StructFormat(w Writer) (n int, err error) {
	if bool(f) {
		return w.Write([]byte("true"))
	}
	return w.Write([]byte("false"))
}
