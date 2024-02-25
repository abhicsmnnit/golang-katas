package main

import "fmt"

type IntegerFamily interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Double[T IntegerFamily](t T) T {
	return 2 * t
}

func main() {
	type MyInt int

	fmt.Println(Double(5))
	fmt.Println(Double(5.2))
	fmt.Println(Double(MyInt(25)))
}
