# Golang structural data formatter

## `import "gopkg.in/structformat.v0"`

![Go Import](https://img.shields.io/badge/import-gopkg.in/structformat.v0-9cf?logo=go&style=for-the-badge)
[![Go Reference](https://img.shields.io/badge/reference-go.dev-007d9c?logo=go&style=for-the-badge)](https://pkg.go.dev/gopkg.in/structformat.v0)

This module gives you greater control over how to format structual data instead of using Golang's default `Printf`
formatting mechanism. It supports outputing to `io.Writer` interface so large data can be formatted in streaming fashion
instead of returning a large string to reduce memory footprint.

## Usage

To support structual formatting for your data types, implement the `structformat.Formatter` interface.

```go
import "gopkg.in/structformat.v0"

type Person struct {
    FirstName string
    LastName string
    Hobbies []string
}

func (p Person) StructFormat(w structformat.Writer) (n int, err error) {
    // Describe your structual data using structformat's built-in formatters
    f := structformat.Struct("Person",
        structformat.KV("Full Name", structformat.Printf("%s %s", p.FirstName, p.LastName)))
    hobbies := structformat.Bracket()
    for _, hobby := range p.Hobbies {
        hobbies.Items = append(hobbies.Items, structformat.String(hobby))
    }
    f.Items = append(f.Items, structformat.KV("Hobbies", hobbies))
    return f.StructFormat(w)
}

func main() {
    person := &Person{
        FirstName: "Adam",
        LastName: "Brody",
        Hobbies: []string{
            "Movies",
            "Singing",
            "Surfing",
        },
    }
    // Format your structual data using structformat.FormatString() or structformat.Format()
    formatted, err := structformat.FormatString(person)
    if err != nil {
        panic(err)
    }
    fmt.Println(formatted)
}
```

And you get the following formatted output:

```
Person {
    Full Name: Adam Brody
    Hobbies: [
        Movies
        Singing
        Surfing
    ]
}
```
