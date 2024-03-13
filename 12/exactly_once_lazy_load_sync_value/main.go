package main

import (
	"fmt"
	"sync"
)

func main() {
	// Calling Parse method multiple times will initialize the parser only once.
	// This can be confirmed by verifying that "Initializing..." is printed only once when this program is executed.
	fmt.Println(Parse("Abhinav"))
	fmt.Println(Parse("Verma"))
	fmt.Println(Parse(""))
}

type Parser interface {
	Parse(string) string
}

var initParserCached func() Parser = sync.OnceValue(initParser)

func Parse(dataToParse string) string {
	parser := initParserCached()
	return parser.Parse(dataToParse)
}

func initParser() Parser {
	// do all sorts of setup and loading here
	fmt.Println("Initializing...")
	return SloooooooooowInitParser{}
}

type SloooooooooowInitParser struct {
}

func (s SloooooooooowInitParser) Parse(in string) string {
	if len(in) > 1 {
		return in[0:1]
	}
	return ""
}
