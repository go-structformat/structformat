package structformat

import "testing"

func TestString(t *testing.T) {
	input := "\nfoo\n\nbar\n"
	expect := input
	actual, err := FormatString(String(input))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}

func TestIndentedString(t *testing.T) {
	input := "\nfoo\n\n\nbar\n\n"
	expect := "\n....foo\n\n\n....bar\n\n"
	actual, err := FormatString(String(input), WithIndent(".."), WithIndentLevel(2))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}
