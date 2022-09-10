package structformat

type StringFormatter []string

var _ Formatter = (*StringFormatter)(nil)

func String(strings ...string) StringFormatter {
	return StringFormatter(strings)
}

func (f StringFormatter) StructFormat() (ret NestedLines) {
	for _, line := range f {
		ret = append(ret, line)
	}
	return
}
