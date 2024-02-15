package main

import (
	"fmt"
)

/*
Write a program that defines a string variable called message with the value "Hi 👩 and 👨" and prints the fourth rune in it as a character, not a number.
*/

func main() {
	s := "Hi 👩 and 👨"
	r := []rune(s)

	fmt.Println("4th rune", string(r[3]))
}
