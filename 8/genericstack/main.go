package main

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

// Can't implement Contains method here since `any` values can't be compared
// Checkout the comparablestack implementation in ../comparablestack

func main() {
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)

	v, ok := s.Pop()
	fmt.Println(v, ok)

	v1, ok := s.Pop()
	fmt.Println(v1, ok)

	v2, ok := s.Pop()
	fmt.Println(v2, ok)

	v3, ok := s.Pop()
	fmt.Println(v3, ok)

	v4, ok := s.Pop()
	fmt.Println(v4, ok)
}
