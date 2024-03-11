package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	res := processAndGather(
		ch,
		func(i int) int {
			return i * 100
		},
		3)
	fmt.Println(res)
}

func processAndGather(in <-chan int, processor func(int) int, concurrency int) []int {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	out := make(chan int)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			for v := range in {
				out <- processor(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var res []int
	for v := range out {
		res = append(res, v)
	}
	return res
}
