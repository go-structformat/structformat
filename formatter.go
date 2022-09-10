package structformat

import (
	"bytes"
	"io"

	"gopkg.in/option.v0"
)

// Implement this interface in your data types to support sturctual formatting.
type Formatter interface {
	// Format the structual data and write to the Writer interface. Returns the number of bytes written in `n`.
	StructFormat(Writer) (n int, err error)
}

// Format a structual data and output to the provided writer.
func Format(writer io.Writer, formatter Formatter, options ...FormatOption) (n int, err error) {
	opts := option.New(options, initWriter(writer), WithIndent(DefaultIndent))
	w := &indentWriter{*opts}
	if n, err = formatter.StructFormat(w); err != nil {
		return
	}
	if opts.endWithNewLine {
		if _, err = writer.Write([]byte{'\n'}); err != nil {
			return
		}
		n += 1
	}
	return
}

// Format a structual data and return the formatted string.
func FormatString(formatter Formatter, options ...FormatOption) (s string, err error) {
	var w bytes.Buffer
	if _, err = Format(&w, formatter, options...); err != nil {
		return
	}
	s = w.String()
	return
}
