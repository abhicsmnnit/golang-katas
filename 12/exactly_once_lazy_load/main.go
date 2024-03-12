package main

import (
	"fmt"
	"sync"
)

type Parser interface {
	Parse(s string) string
}

var parser Parser
var once sync.Once

func main() {
	// Calling Parse method multiple times will initialize the parser only once.
	// This can be confirmed by verifying that "Initializing..." is printed only once when this program is executed.
	fmt.Println(Parse("Abhinav"))
	fmt.Println(Parse("Verma"))
	fmt.Println(Parse(""))
}

func Parse(in string) string {
	once.Do(func() {
		parser = initParser()
	})

	return parser.Parse(in)
}

func initParser() Parser {
	fmt.Println("Initializing...")
	return SlooooowParser{}
}

type SlooooowParser struct {
}

func (sp SlooooowParser) Parse(in string) string {
	if len(in) > 0 {
		return in[0:1]
	}
	return ""
}
