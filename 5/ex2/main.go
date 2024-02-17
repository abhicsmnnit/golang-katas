package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Filename not provided")
	}

	filename := os.Args[1]

	len, err := fileLen(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)
}

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	len := 0
	bytes := make([]byte, 2048)
	for {
		count, err := f.Read(bytes)
		len += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}

	return len, nil
}
