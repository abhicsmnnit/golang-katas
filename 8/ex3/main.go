package main

import (
	"fmt"
)

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	// f    func(t1, t2 T) bool
}

// adds a new element to the end of the linked list
func (ll *LinkedList[T]) Add(val T) {
	node := &Node[T]{val: val}

	if ll.head == nil {
		ll.head = node
	} else {
		ll.tail.next = node
	}

	ll.tail = node
}

// adds an element at the specified position in the linked list
func (ll *LinkedList[T]) Insert(val T, pos int) {
	node := &Node[T]{val: val}

	if ll.head == nil {
		ll.head = node
		ll.tail = node
		return
	}

	if pos <= 0 {
		node.next = ll.head
		ll.head = node
		return
	}

	cur := ll.head
	for i := 1; i < pos; i++ {
		if cur.next == nil {
			cur.next = node
			ll.tail = node
			return
		}
		cur = cur.next
	}

	node.next = cur.next
	cur.next = node
}

// returns the position of the supplied value, -1 if it's not present
func (ll *LinkedList[T]) Index(val T) int {
	// fmt.Println("Index", val)
	pos := 0

	for node := ll.head; node != nil; node = node.next {
		// fmt.Println("node val", node.val)
		if node.val == val {
			return pos
		}
		pos++
	}

	return -1
}

// func NewLinkedList[T any](f func(t1, t2 T) bool) LinkedList[T] {
// 	return LinkedList[T]{f: f}
// }

type Node[T any] struct {
	val  T
	next *Node[T]
}

func main() {
	lli := LinkedList[int]{}
	lli.Add(1)
	lli.Add(2)
	lli.Add(3)
	fmt.Println(lli.Index(1))
	fmt.Println(lli.Index(2))
	fmt.Println(lli.Index(3))
	fmt.Println(lli.Index(4))
	fmt.Println(lli.Index(5))

	lli.Insert(4, 1)

	fmt.Println(lli.Index(1))
	fmt.Println(lli.Index(2))
	fmt.Println(lli.Index(3))
	fmt.Println(lli.Index(4))
	fmt.Println(lli.Index(5))

	lli.Insert(5, 0)
	lli.Insert(6, 10)

	fmt.Println(lli.Index(1))
	fmt.Println(lli.Index(2))
	fmt.Println(lli.Index(3))
	fmt.Println(lli.Index(4))
	fmt.Println(lli.Index(5))
	fmt.Println(lli.Index(6))
}
