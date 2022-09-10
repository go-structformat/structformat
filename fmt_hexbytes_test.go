package structformat

import "testing"

func TestHexBytes(t *testing.T) {
	input := []byte{0xde, 0xad, 0xbe, 0xef, 0xab, 0xcd, 0xef}
	expect := "0xdeadbeefabcdef"
	actual, err := FormatString(HexBytes(input))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}
