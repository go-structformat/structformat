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

func (f *BlockFormatter) StructFormat() NestedLines {
	var sublevel NestedLines
	for _, item := range f.Items {
		sublevel = append(sublevel, item.StructFormat()...)
	}
	return NestedLines{f.Open, sublevel, f.Close}
}
