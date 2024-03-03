package main

import (
	"fmt"

	mymath "github.com/abhicsmnnit/ch10exercise/math"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Printf(" %p %p\n", &i, &v) // Different variables in go1.22 and later; same before that.
	}

	fmt.Println(mymath.Add(10, 10))
}
