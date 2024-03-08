package main

import (
	"context"
	"fmt"
)

func main() {
	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// result := processConcurrently(x)
	// fmt.Println(result)

	// // This works!
	// for v := range countTo(10) {
	// 	fmt.Println(v)
	// }

	// // This blocks the goroutine countTo because we exits the loop early and countTo is waiting for a value to be read from the channel
	// for v := range countTo(10) {
	// 	if v > 5 {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	// This communicates to the goroutine in countToWithCtx via context, so it can clean up and avoid "leaking"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for v := range countToWithCtx(10, ctx) {
		if v > 5 {
			break
		}
		fmt.Println(v)
	}
}

func countToWithCtx(num int, ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func countTo(num int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < num; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
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
