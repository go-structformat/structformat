package structformat

import (
	"encoding/hex"
)

type HexdumpFormatter []byte

var _ Formatter = (*HexdumpFormatter)(nil)

// Returns a HexdumpFormatter that format the bytes data using hex.Dumper().
func Hexdump(data []byte) HexdumpFormatter {
	return HexdumpFormatter(data)
}

func (f HexdumpFormatter) StructFormat(w Writer) (n int, err error) {
	return hex.Dumper(w.RemoveTrailingNewLines()).Write(f)
}
