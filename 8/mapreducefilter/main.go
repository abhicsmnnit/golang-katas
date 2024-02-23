package main

import "fmt"

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	t2s := make([]T2, len(s))

	for i, v := range s {
		t2s[i] = f(v)
	}

	return t2s
}

func main() {
	fmt.Println(Map([]string{"Go", "Java", "Rust", "Python"}, func(s string) int {
		return len(s)
	}))
}
