package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Not a zip file!")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Invalid zip file")
	}
}
