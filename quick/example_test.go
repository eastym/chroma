package quick_test

import (
	"log"
	"os"

	"github.com/eastym/chroma/v2/quick"
)

func Example() {
	code := `package main

func main() { }
`
	err := quick.Highlight(os.Stdout, code, "go", "html", "monokai")
	if err != nil {
		log.Fatal(err)
	}
}
