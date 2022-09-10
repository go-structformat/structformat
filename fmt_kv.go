package structformat

type KVFormatter struct {
	Name            string
	Value           Formatter
	NewLineForValue bool
}

var _ Formatter = (*KVFormatter)(nil)

func KV(name string, value Formatter, newLineForValue bool) *KVFormatter {
	return &KVFormatter{Name: name, Value: value, NewLineForValue: newLineForValue}
}

func (f *KVFormatter) StructFormat() NestedLines {
	firstline := f.Name + `: `
	value := f.Value.StructFormat()
	if len(value) == 0 {
		firstline += `( EMPTY )`
		return NestedLines{firstline}
	} else if str, ok := value[0].(string); ok && !f.NewLineForValue {
		firstline += str
		return append(NestedLines{firstline}, value[1:]...)
	}
	return NestedLines{firstline + `{`, value, `}`}
}
