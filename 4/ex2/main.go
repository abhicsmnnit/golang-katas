package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ints := make([]int, 100)

	for i := 0; i < 100; i++ {
		ints[i] = rand.Intn(100)
	}

	for _, v := range ints {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")

		case v%2 == 0:
			fmt.Println("Two!")

		case v%3 == 0:
			fmt.Println("Three!")

		default:
			fmt.Println("Never mind")
		}
	}
}
