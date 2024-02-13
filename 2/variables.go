package main

import "fmt"

func main() {
	var b byte = 255
	fmt.Println(b)
	fmt.Println(b + 1) // Wraps around to 0

	var smallI int32 = 2_147_483_647
	fmt.Println(smallI)
	fmt.Println(smallI + 1) // Wraps around to -2_147_483_648

	var bigI uint64 = 18_446_744_073_709_551_615
	fmt.Println(bigI)
	fmt.Println(bigI + 1) // Wraps around to 0
}
