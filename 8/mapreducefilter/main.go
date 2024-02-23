package main

import "fmt"

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	t2s := make([]T2, len(s))

	for i, v := range s {
		t2s[i] = f(v)
	}

	return t2s
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer

	for _, v := range s {
		r = f(r, v)
	}

	return r
}

func Filter[T any](s []T, f func(t T) bool) []T {
	filteredSlice := []T{}
	for _, v := range s {
		if f(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func main() {
	lengths := Map([]string{"Go", "Java", "Rust", "Python"}, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)

	sum := Reduce([]int{1, 2, 3, 4, 5}, 0, func(a int, b int) int {
		return a + b
	})
	fmt.Println(sum)

	evenNumbers := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(evenNumbers)
}
