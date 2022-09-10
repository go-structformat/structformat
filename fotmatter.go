package structformat

import (
	"strings"
)

// interface{} being either string or []string
type NestedLines []interface{}

type Formatter interface {
	StructFormat() NestedLines
}

func Format(formatter Formatter, indent string) string {
	return strings.Join(flatten(formatter.StructFormat(), 0, indent), "\n")
}

func flatten(nestedLines NestedLines, level int, indent string) (lines []string) {
	for _, line := range nestedLines {
		if line == nil {
			continue
		} else if str, ok := line.(string); ok {
			lines = append(lines, strings.Repeat(indent, level)+str)
		} else if arr, ok := line.(NestedLines); ok {
			lines = append(lines, flatten(arr, level+1, indent)...)
		}
	}
	return
}
