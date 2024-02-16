package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i // total shadowing the total variable outside the loop. Instead of assign, we can use = operator.
		fmt.Println(total)
	}

	fmt.Println(total)
}
