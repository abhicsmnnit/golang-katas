package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json: "title"`
	Author string `json: author`
	Price  int    `json: "price"`
}

func main() {
	toFile := Book{
		Title:  "Learning Go",
		Author: "Jon Bodner",
		Price:  99,
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "book-sample")

	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}

	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	var fromFile Book
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}

	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", fromFile)
}
