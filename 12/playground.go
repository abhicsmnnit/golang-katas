package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := processConcurrently(x)
	fmt.Println(result)
}

func processConcurrently(input []int) []int {
	// numGoroutines := 1
	numGoroutines := 5
	in := make(chan int, numGoroutines)
	out := make(chan int, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for val := range in {
				result := process(val)
				out <- result
			}
		}()
	}

	go func() {
		for _, v := range input {
			in <- v
		}
		close(in)
	}()

	output := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = <-out
	}

	return output
}

func process(val int) int {
	return 2 * val
}
