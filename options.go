package structformat

import "io"

type FormatOption func(*indentWriter)

func withWriter(writer io.Writer) FormatOption {
	return func(w *indentWriter) {
		w.Writer = writer
	}
}

const DefaultIndent = "    "

// Use the provided string for each indent level
func WithIndent(indent string) FormatOption {
	return func(w *indentWriter) {
		w.indent = indent
	}
}

// Use the provided indent level
func WithIndentLevel(level int) FormatOption {
	return func(w *indentWriter) {
		w.level = level
	}
}
