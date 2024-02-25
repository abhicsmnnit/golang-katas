package main

import (
	"cmp"
	"fmt"
)

type Tree[T any] struct {
	root *Node[T]
	f    OrderableFunc[T]
}

func (t *Tree[T]) Add(val T) {
	t.root = t.root.Add(val, t.f)
}

func (t *Tree[T]) Contains(val T) bool {
	return t.root.Contains(val, t.f)
}

func NewTree[T any](f OrderableFunc[T]) Tree[T] {
	return Tree[T]{
		f: f,
	}
}

type OrderableFunc[T any] func(t1, t2 T) int

type Node[T any] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func (n *Node[T]) Add(v T, f OrderableFunc[T]) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}

	switch c := f(v, n.val); {
	case c < 0:
		n.left = n.left.Add(v, f)
	case c > 0:
		n.right = n.right.Add(v, f)
	}

	return n
}

func (n *Node[T]) Contains(v T, f OrderableFunc[T]) bool {
	if n == nil {
		return false
	}

	switch c := f(v, n.val); {
	case c < 0:
		return n.left.Contains(v, f)
	case c > 0:
		return n.right.Contains(v, f)
	default:
		return true
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) Order(other Person) int {
	c := cmp.Compare(p.name, other.name)
	if c == 0 {
		c = cmp.Compare(p.age, other.age)
	}
	return c
}

func main() {
	tree := NewTree(func(i1, i2 int) int {
		return cmp.Compare(i1, i2)
	})
	tree.Add(5)
	fmt.Println(tree.Contains(5))
	fmt.Println(tree.Contains(10))
	tree.Add(10)
	fmt.Println(tree.Contains(10))

	personTree := NewTree(Person.Order)
	personTree.Add(Person{"Abhinav", 30})
	personTree.Add(Person{"Siddhanta", 35})
	personTree.Add(Person{"Siddhanta", 33})
	fmt.Println(personTree.Contains(Person{"Siddhanta", 35}))
	fmt.Println(personTree.Contains(Person{"Abhinav", 40}))
	fmt.Println(personTree.Contains(Person{"Siddhanta", 33}))
}
