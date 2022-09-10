package structformat

import (
	"bytes"
	"io"
)

type Writer interface {
	io.Writer

	IndentString() string
	LineNum() int
	LineOffset() int
	NonIndentLineOffset() int
	IsEmpty() bool

	// Return a new writer with indent level set to be current
	// indent level plus levelDelta. If levelDelta is 0, indent
	// is reset to empty level.
	Indent(levelDelta int) Writer
	EmptyCheck(prependIfNotEmpty string) Writer
	RemoveTrailingNewLines() Writer
}

type indentWriter struct {
	formatOptions
}

var _ io.Writer = (*indentWriter)(nil)

func (w *indentWriter) Write(data []byte) (n int, err error) {
	for len(data) > 0 {
		var i int
		for i = 0; i < len(data) && data[i] != '\n'; i++ {
		}
		if i > 0 {
			if w.lineOffset == 0 {
				if w.level > 0 {
					for l := 0; l < w.level; l++ {
						if _, err = w.writer.Write([]byte(w.indent)); err != nil {
							return
						}
					}
					n += w.level * len(w.indent)
					w.lineOffset += w.level * len(w.indent)
				}
				w.lineNum += 1
			}
			if _, err = w.writer.Write(data[:i]); err != nil {
				return
			}
			n += i
			w.lineOffset += i
		}
		if i < len(data) && data[i] == '\n' {
			if w.lineOffset == 0 {
				w.lineNum += 1
			} else {
				w.lineOffset = 0
			}
			var j int
			for j = i + 1; j < len(data) && data[j] == '\n'; j++ {
			}
			if _, err = w.writer.Write(data[i:j]); err != nil {
				return
			}
			n += j - i
			w.lineNum += j - i - 1
			data = data[j:]
		} else {
			break
		}
	}
	return
}

func (w indentWriter) IndentString() string {
	return w.indent
}

func (w indentWriter) LineNum() int {
	return w.lineNum
}

func (w indentWriter) LineOffset() int {
	return w.lineOffset
}

func (w indentWriter) NonIndentLineOffset() int {
	offset := w.lineOffset - w.level*len(w.indent)
	if offset <= 0 {
		return 0
	}
	return offset
}

func (w indentWriter) IsEmpty() bool {
	return w.lineNum > 0
}

func (w indentWriter) Indent(levelDelta int) Writer {
	level := w.level + levelDelta
	if levelDelta == 0 {
		level = 0
	}
	writer := w
	writer.level = level
	return &writer
}

func (w indentWriter) EmptyCheck(prependIfNotEmpty string) Writer {
	return &emptyCheckWriter{
		Writer:            &w,
		prependIfNotEmpty: prependIfNotEmpty,
	}
}

func (w indentWriter) RemoveTrailingNewLines() Writer {
	return &removeTrailingNewLinesWriter{&w, 0}
}

type emptyCheckWriter struct {
	Writer
	prependIfNotEmpty string
	notEmpty          bool
}

func (w *emptyCheckWriter) Write(data []byte) (n int, err error) {
	var i int
	if !w.notEmpty && len(data) > 0 {
		if i, err = w.Writer.Write([]byte(w.prependIfNotEmpty)); err != nil {
			return
		}
		n += i
		w.notEmpty = true
	}
	if i, err = w.Writer.Write(data); err != nil {
		return
	}
	n += i
	return
}

func (w emptyCheckWriter) IsEmpty() bool {
	return !w.notEmpty
}

func (w emptyCheckWriter) Indent(levelDelta int) Writer {
	writer := w
	writer.Writer = w.Writer.Indent(levelDelta)
	return &writer
}

func (w emptyCheckWriter) EmptyCheck(prependIfNotEmpty string) Writer {
	writer := w
	if w.notEmpty {
		writer.notEmpty = false
		writer.prependIfNotEmpty = prependIfNotEmpty
	} else {
		writer.prependIfNotEmpty += prependIfNotEmpty
	}
	return &writer
}

func (w emptyCheckWriter) RemoveTrailingNewLines() Writer {
	return &removeTrailingNewLinesWriter{&w, 0}
}

type removeTrailingNewLinesWriter struct {
	Writer
	bufferedNewLines int
}

func (w *removeTrailingNewLinesWriter) Write(data []byte) (n int, err error) {
	for len(data) > 0 {
		var i int
		for i = 0; i < len(data) && data[i] != '\n'; i++ {
		}
		if i > 0 {
			if w.bufferedNewLines > 0 {
				if _, err = w.Writer.Write(bytes.Repeat([]byte{'\n'}, w.bufferedNewLines)); err != nil {
					return
				}
				n += w.bufferedNewLines
				w.bufferedNewLines = 0
			}
			if _, err = w.Writer.Write(data[:i]); err != nil {
				return
			}
			n += i
		}
		if i < len(data) && data[i] == '\n' {
			var j int
			for j = i + 1; j < len(data) && data[j] == '\n'; j++ {
			}
			w.bufferedNewLines += j - i
			data = data[j:]
		} else {
			break
		}
	}
	return
}

func (w removeTrailingNewLinesWriter) Indent(levelDelta int) Writer {
	writer := w
	writer.Writer = w.Writer.Indent(levelDelta)
	return &writer
}

func (w removeTrailingNewLinesWriter) EmptyCheck(prependIfNotEmpty string) Writer {
	writer := w
	writer.Writer = w.Writer.EmptyCheck(prependIfNotEmpty)
	return &writer
}

func (w removeTrailingNewLinesWriter) RemoveTrailingNewLines() Writer {
	return &w
}
