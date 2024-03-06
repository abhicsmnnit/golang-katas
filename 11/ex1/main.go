package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed english_rights.txt
var englishRights string

//go:embed chinese_rights.txt
var chineseRights string

func main() {
	var language string
	if len(os.Args) < 2 {
		language = "english"
	} else {
		language = os.Args[1]
	}

	switch strings.ToLower(language) {
	case "english":
		fmt.Println(englishRights)
	case "chinese":
		fmt.Println(chineseRights)
	default:
		fmt.Println("language not supported:", language)
	}
}
