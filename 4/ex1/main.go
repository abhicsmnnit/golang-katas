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

	fmt.Println(len(ints), ints)
}
