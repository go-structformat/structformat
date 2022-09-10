package structformat

import (
	"bufio"
	"strings"
)

type StringFormatter []string

var _ Formatter = (*StringFormatter)(nil)

func String(s string) StringFormatter {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return StringFormatter(lines)
}

func (f StringFormatter) StructFormat() (ret NestedLines) {
	for _, line := range f {
		ret = append(ret, line)
	}
	return
}
