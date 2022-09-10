package structformat

import "encoding/hex"

type HexBytesFormatter []byte

var _ Formatter = (*HexBytesFormatter)(nil)

func HexBytes(data []byte) HexBytesFormatter {
	return HexBytesFormatter(data)
}

func (f HexBytesFormatter) StructFormat() NestedLines {
	return NestedLines{`0x` + hex.EncodeToString(f)}
}
