package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	ch := make(chan int)

	go func() {
		for _, v := range nums {
			ch <- v
		}
	}()

	result := processChannel(ch)
	fmt.Println(result)
}

func processChannel(ch chan int) []int {
	const concurrency = 10

	resultsChan := make(chan int, concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			v := <-ch
			resultsChan <- process(v)
		}()
	}

	var out []int
	for i := 0; i < concurrency; i++ {
		out = append(out, <-resultsChan)
	}
	return out
}

func process(num int) int {
	return 2 * num
}
