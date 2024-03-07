package main

import "fmt"

func main() { // creates a deadlock
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 10
		ch2Val := <-ch2
		fmt.Println("goroutine:", ch2Val)
	}()

	ch2 <- 20
	ch1Val := <-ch1
	fmt.Println("main:", ch1Val)
}
