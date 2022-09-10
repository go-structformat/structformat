package structformat

import (
	"encoding/hex"
	"strings"
)

type HexdumpFormatter []byte

var _ Formatter = (*HexdumpFormatter)(nil)

func Hexdump(data []byte) HexdumpFormatter {
	return HexdumpFormatter(data)
}

func (f HexdumpFormatter) StructFormat() (ret NestedLines) {
	for _, line := range strings.Split(hex.Dump(f), "\n") {
		if strings.TrimSpace(line) != "" {
			ret = append(ret, line)
		}
	}
	return
}
