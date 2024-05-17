package main

import "fmt"

/*
Create a function that launches two goroutines.
Each goroutine writes 10 numbers to its own channel.
Use a for-select loop to read from both channels, printing out the number and the goroutine that wrote the value.
Make sure that your function exits after all values are read and that none of your goroutines leak.
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			ch2 <- i * 100
		}
		close(ch2)
	}()

	count := 0
	for count < 2 {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				count++
				break
			}
			fmt.Println(v, 1)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count++
				break
			}
			fmt.Println(v, 2)
		}
	}
}
