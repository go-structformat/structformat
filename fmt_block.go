package structformat

type BlockFormatter struct {
	Open  string
	Close string
	Items []Formatter
}

var _ Formatter = (*BlockFormatter)(nil)

func Block(items ...Formatter) *BlockFormatter {
	return &BlockFormatter{Open: "{", Close: "}", Items: items}
}

func Bracket(items ...Formatter) *BlockFormatter {
	return &BlockFormatter{Open: "[", Close: "]", Items: items}
}

func (f *BlockFormatter) StructFormat(w Writer) (n int, err error) {
	var i int
	if i, err = w.Write([]byte(f.Open)); err != nil {
		return
	}
	n += i
	subwriter := w.Indent(1).EmptyCheck("\n")
	for _, item := range f.Items {
		if i, err = item.StructFormat(subwriter); err != nil {
			return
		}
		n += i
		if i, err = subwriter.Write([]byte("\n")); err != nil {
			return
		}
		n += i
	}
	i, err = w.Write([]byte(f.Close))
	n += i
	return
}
