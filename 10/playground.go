package main

import (
	"fmt"

	"github.com/abhicsmnnit/ch10exercise/v2/math"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Printf(" %p %p\n", &i, &v) // Different variables in go1.22 and later; same before that.
	}

	fmt.Println(math.Add(10, 10))
	fmt.Println(math.Add(10.01, 10.01))
}
