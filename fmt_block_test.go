package structformat

import "testing"

func TestEmptyBlock(t *testing.T) {
	expect := "{}"
	actual, err := FormatString(Block())
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}

func TestBlock(t *testing.T) {
	expect := `{
    foo
    bar
}`
	actual, err := FormatString(Block(String("foo"), String("bar")))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}
