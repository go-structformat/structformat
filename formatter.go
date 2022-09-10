package structformat

import (
	"strings"
)

// interface{} being either string or []string
type NestedLines []interface{}

type Formatter interface {
	StructFormat() NestedLines
}

func Format(formatter Formatter, options ...FormatOption) string {
	opts := combineOptions(append([]FormatOption{
		FormatWithIndent(DefaultIndent),
	}, options...))
	return strings.Join(flatten(formatter.StructFormat(), 0, opts.indent), "\n")
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

const DefaultIndent = "    "

type FormatOption func(*formatOptions)

type formatOptions struct {
	indent string
}

func FormatWithIndent(indent string) FormatOption {
	return func(o *formatOptions) {
		o.indent = indent
	}
}
