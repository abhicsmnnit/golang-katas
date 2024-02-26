package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// // Using sentinel errors
	// data := []byte("Not a zip file!")
	// notAZipFile := bytes.NewReader(data)
	// _, err := zip.NewReader(notAZipFile, int64(len(data)))
	// if err == zip.ErrFormat {
	// 	fmt.Println("Invalid zip file")
	// }

	err := fileChecker("DefinitelyNotAFile.txt")

	if err != nil {
		fmt.Println(err)
		if wrappedError := errors.Unwrap(err); wrappedError != nil { // Unwrap the error
			fmt.Println(wrappedError)
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err) // Wrap the error
	}
	f.Close()
	return nil
}
