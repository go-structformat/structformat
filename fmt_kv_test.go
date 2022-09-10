package structformat

import "testing"

func TestEmptyKV(t *testing.T) {
	expect := "foo: ( EMPTY )"
	actual, err := FormatString(KV("foo", String(""), WithNewLine()))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}

func TestKVNoLineBreak(t *testing.T) {
	expect := "foo: bar"
	actual, err := FormatString(KV("foo", String("bar")))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}

func TestKVWithLineBreak(t *testing.T) {
	expect := `foo:
    bar
    baz`
	actual, err := FormatString(KV("foo", String("bar\nbaz"), WithNewLine()))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}
