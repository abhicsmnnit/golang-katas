package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Waiting for my 3 friends")

	go func() {
		defer wg.Done()
		rohitIsHere()
	}()
	go func() {
		defer wg.Done()
		saurabhIsHere()
	}()
	go func() {
		defer wg.Done()
		prachiIsHere()
	}()

	wg.Wait()
	fmt.Println("Off we go")
}

func rohitIsHere() {
	fmt.Println("Rohit is here")
}

func saurabhIsHere() {
	fmt.Println("Saurabh is here")
}

func prachiIsHere() {
	fmt.Println("Prachi is here")
}
