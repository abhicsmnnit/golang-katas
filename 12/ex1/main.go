package main

import (
	"fmt"
	"sync"
)

/*
Create a function that launches three goroutines that communicate using a channel.
The first two goroutines each write 10 numbers to the channel.
The third goroutine reads all the numbers from the channel and prints them out.
The function should exit when all values have been printed out.
Make sure that none of the goroutines leak. You can create additional goroutines if needed.
*/

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i * 100
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	fmt.Println("Exiting...")
}
