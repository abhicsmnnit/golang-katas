package main

import "fmt"

/*
Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
Create a subslice containing the first two values; a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values.
Print out all four slices.
*/

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	fmt.Println(greetings[:2])
	fmt.Println(greetings[1:4])
	fmt.Println(greetings[3:])
	fmt.Println(greetings)
}
