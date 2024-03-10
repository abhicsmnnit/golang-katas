package main

import "fmt"

func main() {
	in1 := make(chan int)
	in2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i += 2 {
			in1 <- i
		}
		close(in1)
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			in2 <- i
		}
		close(in2)
	}()

	result := readFromTwoChannels(in1, in2)
	fmt.Println(result)
}

func readFromTwoChannels(in1, in2 chan int) []int {
	var out []int

	for count := 0; count < 2; {
		select {
		case v, ok := <-in1:
			if !ok {
				in1 = nil // the case will never succeed again!
				count++
				continue
			}
			// process the v that was read from in1
			out = append(out, v)
		case v, ok := <-in2:
			if !ok {
				in2 = nil // the case will never succeed again!
				count++
				continue
			}
			// process the v that was read from in2
			out = append(out, v)
		}
	}

	return out
}
