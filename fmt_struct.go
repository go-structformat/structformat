package structformat

import "fmt"

type StructFormatter struct {
	Name  string
	Items []Formatter
}

var _ Formatter = (*StructFormatter)(nil)

// Return a StructFormatter that has a name and arbitrary number of items.
func Struct(name string, items ...Formatter) *StructFormatter {
	return &StructFormatter{Name: name, Items: items}
}

func (f *StructFormatter) StructFormat(w Writer) (n int, err error) {
	var i int
	if i, err = fmt.Fprintf(w, "%s {", f.Name); err != nil {
		return
	}
	n += i
	subwriter := w.EmptyCheck("\n").Indent(1)
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
	if i, err = w.Write([]byte("}")); err != nil {
		return
	}
	n += i
	return
}
