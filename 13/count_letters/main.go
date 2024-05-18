package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	letterMap := map[string]int{}
	buffer := make([]byte, 2048)

	for {
		n, err := r.Read(buffer)

		for _, b := range buffer[:n] {
			if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
				letterMap[string(b)]++
			}
		}

		if err == io.EOF {
			return letterMap, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func main() {
	// Counting letters with a string reader
	counter, err := countLetters(getStringReader("The quick brown fox jumps over a lazy dog"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(counter)

	// Counting letters with a gzip file reader
	gr, closer, err := getGZipReader("test.gzip.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closer()
	counter, err = countLetters(gr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(counter)
}

func getStringReader(s string) io.Reader {
	return strings.NewReader(s)
}

func getGZipReader(filename string) (*gzip.Reader, func(), error) {
	fr, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	gr, err := gzip.NewReader(fr)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return gr, func() {
		gr.Close()
		fr.Close()
	}, nil
}
