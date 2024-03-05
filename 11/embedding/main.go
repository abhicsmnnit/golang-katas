package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed common_passwords.txt
var commonPasswordsFile string

func main() {
	if len(os.Args) < 2 {
		log.Fatal("password missing")
	}

	commonPasswords := strings.Split(commonPasswordsFile, ",")

	for _, v := range commonPasswords {
		if v == os.Args[1] {
			fmt.Println("true")
			os.Exit(0)
		}
	}

	fmt.Println("false")
}
