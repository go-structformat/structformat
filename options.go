package structformat

import "io"

type FormatOption func(*formatOptions)

type formatOptions struct {
	*writerContext
	writer         io.Writer
	indent         string
	level          int
	endWithNewLine bool
}

type writerContext struct {
	lineNum    int
	lineOffset int
}

func initWriter(writer io.Writer) FormatOption {
	return func(o *formatOptions) {
		o.writer = writer
		o.writerContext = &writerContext{}
	}
}

const DefaultIndent = "    "

// Use the provided string for each indent level
func WithIndent(indent string) FormatOption {
	return func(o *formatOptions) {
		o.indent = indent
	}
}

// Use the provided indent level
func WithIndentLevel(level int) FormatOption {
	return func(o *formatOptions) {
		o.level = level
	}
}

func EndWithNewLine() FormatOption {
	return func(o *formatOptions) {
		o.endWithNewLine = true
	}
}
