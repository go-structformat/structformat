package structformat

import "encoding/hex"

type HexBytesFormatter []byte

var _ Formatter = (*HexBytesFormatter)(nil)

// Return a HexBytesFormatter that formats the bytes data using hex.Encoder.
func HexBytes(data []byte) HexBytesFormatter {
	return HexBytesFormatter(data)
}

func (f HexBytesFormatter) StructFormat(w Writer) (n int, err error) {
	if len(f) == 0 {
		return
	}
	var i int
	if i, err = w.Write([]byte("0x")); err != nil {
		return
	}
	n += i
	if i, err = hex.NewEncoder(w).Write(f); err != nil {
		return
	}
	n += i
	return
}
