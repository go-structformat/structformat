package structformat

import (
	"testing"
)

type Person struct {
	FirstName string
	LastName  string
	Hobbies   []string
}

func (p Person) StructFormat(w Writer) (n int, err error) {
	// Describe your structual data using structformat's built-in formatters
	f := Struct("Person",
		KV("Full Name", Printf("%s %s", p.FirstName, p.LastName)))
	hobbies := Bracket()
	for _, hobby := range p.Hobbies {
		hobbies.Items = append(hobbies.Items, String(hobby))
	}
	f.Items = append(f.Items, KV("Hobbies", hobbies))
	return f.StructFormat(w)
}

func TestExample(t *testing.T) {
	expect := `Person {
    Full Name: Adam Brody
    Hobbies: [
        Movies
        Singing
        Surfing
    ]
}`
	person := &Person{
		FirstName: "Adam",
		LastName:  "Brody",
		Hobbies: []string{
			"Movies",
			"Singing",
			"Surfing",
		},
	}
	// Format your structual data using FormatString() or Format()
	actual, err := FormatString(person)
	if err != nil {
		panic(err)
	}
	if actual != expect {
		t.Errorf("expects %#v but got %#v", expect, actual)
	}
}
