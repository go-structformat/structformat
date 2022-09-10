package structformat

import (
	"fmt"

	"gopkg.in/option.v0"
)

type KVFormatter struct {
	Name            string
	Value           Formatter
	NewLineForValue bool
}

var _ Formatter = (*KVFormatter)(nil)

// Return a KVFormatter represents a key-value pair.
// Oftenly used inside the StructFormatter for form a map like structure.
func KV(name string, value Formatter, options ...KVOption) *KVFormatter {
	kv := option.New(options)
	kv.Name = name
	kv.Value = value
	return kv
}

func (f *KVFormatter) StructFormat(w Writer) (n int, err error) {
	var i int
	if i, err = fmt.Fprintf(w, "%s:", f.Name); err != nil {
		return
	}
	n += i
	subwriter := w.EmptyCheck(" ")
	if f.NewLineForValue {
		subwriter = w.EmptyCheck("\n").Indent(1)
	}
	if i, err = f.Value.StructFormat(subwriter); err != nil {
		return
	}
	n += i
	if i == 0 {
		i, err = fmt.Fprintf(w, " ( EMPTY )")
		n += i
	}
	return
}

type KVOption func(*KVFormatter)

func KVNewLine() KVOption {
	return func(k *KVFormatter) {
		k.NewLineForValue = true
	}
}
