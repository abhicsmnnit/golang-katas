package main

import "fmt"

type IntBST struct {
	val         int
	left, right *IntBST
}

func (ib *IntBST) Insert(val int) *IntBST {
	if ib == nil {
		return &IntBST{val: val}
	}

	if val < ib.val {
		ib.left = ib.left.Insert(val)
	} else if val > ib.val {
		ib.right = ib.right.Insert(val)
	}

	return ib
}

func (ib *IntBST) Contains(val int) bool {
	if ib == nil {
		return false
	}

	switch {
	case ib == nil:
		return false
	case val < ib.val:
		return ib.left.Contains(val)
	case val > ib.val:
		return ib.right.Contains(val)
	default:
		return true
	}
}

func main() {
	var ib *IntBST
	ib = ib.
		Insert(5).
		Insert(3).
		Insert(7).
		Insert(4).
		Insert(2).
		Insert(6).
		Insert(8)

	fmt.Println(ib.Contains(6), ib.Contains(10))
}
