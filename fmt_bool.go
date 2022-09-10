package structformat

type BoolFormatter bool

var _ Formatter = (*BoolFormatter)(nil)

func Bool(value bool) BoolFormatter {
	return BoolFormatter(value)
}

func (f BoolFormatter) StructFormat() (ret NestedLines) {
	if bool(f) {
		return NestedLines{`true`}
	}
	return NestedLines{`false`}
}
