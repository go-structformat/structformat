package structformat

import "testing"

func TestStructSimpleString(t *testing.T) {
	expect := `Test {
    foo
    bar
}`
	actual, err := FormatString(Struct("Test", String("foo"), String("bar")))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}

func TestStructKV(t *testing.T) {
	expect := `Test {
    foo {
        abc: def
    }
}`
	actual, err := FormatString(Struct("Test", Struct("foo", KV("abc", String("def")))))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}
