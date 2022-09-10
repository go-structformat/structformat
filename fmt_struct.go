package structformat

type StructFormatter struct {
	Name  string
	Items []Formatter
}

var _ Formatter = (*StructFormatter)(nil)

func Struct(name string, items ...Formatter) *StructFormatter {
	return &StructFormatter{Name: name, Items: items}
}

func (f *StructFormatter) StructFormat() NestedLines {
	var sublevel NestedLines
	for _, item := range f.Items {
		sublevel = append(sublevel, item.StructFormat()...)
	}
	return NestedLines{
		f.Name + ` {`,
		sublevel,
		`}`,
	}
}
