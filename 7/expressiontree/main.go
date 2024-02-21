package main

import (
	"errors"
	"fmt"
	"log"
)

type exprTreeNode struct {
	val    treeVal
	lChild *exprTreeNode
	rChild *exprTreeNode
}

func (node *exprTreeNode) evaluate() (int, error) {
	if node == nil {
		return 0, errors.New("can't evaluate nil node")
	}

	switch val := node.val.(type) {
	case nil:
		return 0, errors.New("invalid node: nil")
	case number:
		return int(val), nil
	case operator:
		lVal, err := node.lChild.evaluate()
		if err != nil {
			return 0, err
		}

		rVal, err := node.rChild.evaluate()
		if err != nil {
			return 0, err
		}

		return val(lVal, rVal), nil
	default:
	}
	return 0, nil
}

// treeVal defines an unexported marker interface that makes it clear
// which types can be assigned to val in exprTreeNode
type treeVal interface {
	isToken()
}

type number int

func (n number) isToken() {

}

type operator func(int, int) int

func (o operator) isToken() {

}

func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

var operators = map[string]operator{
	"+": func(n1, n2 int) int { return n1 + n2 },
	"-": func(n1, n2 int) int { return n1 - n2 },
	"*": func(n1, n2 int) int { return n1 * n2 },
	"/": func(n1, n2 int) int { return n1 / n2 },
}

func main() {
	exprTreeNode, err := parse("10*20+500")

	if err != nil {
		log.Fatal("Invalid expression")
	}

	res, err := exprTreeNode.evaluate()
	fmt.Println(res, err)
}

func parse(exprString string) (*exprTreeNode, error) {
	// returning a hard-coded value

	return validTree()
	// return invalidTree1()
	// return invalidTree2()
	// return invalidTree3()
}

func validTree() (*exprTreeNode, error) {
	return &exprTreeNode{
		val: operators["+"],
		lChild: &exprTreeNode{
			val:    operators["*"],
			lChild: &exprTreeNode{val: number(10)},
			rChild: &exprTreeNode{val: number(20)},
		},
		rChild: &exprTreeNode{val: number(500)},
	}, nil
}

func invalidTree1() (*exprTreeNode, error) {
	return &exprTreeNode{
		val: operators["+"],
		lChild: &exprTreeNode{
			val:    operators["*"],
			lChild: &exprTreeNode{val: number(10)},
			// rChild: &exprTreeNode{val: number(20)},
		},
		rChild: &exprTreeNode{val: number(500)},
	}, nil
}

func invalidTree2() (*exprTreeNode, error) {
	return &exprTreeNode{
		val: operators["+"],
		lChild: &exprTreeNode{
			// val:    operators["*"],
			lChild: &exprTreeNode{val: number(10)},
			rChild: &exprTreeNode{val: number(20)},
		},
		rChild: &exprTreeNode{val: number(500)},
	}, nil
}

func invalidTree3() (*exprTreeNode, error) {
	return &exprTreeNode{
		val: operators["+"],
		lChild: &exprTreeNode{
			val:    operators["*"],
			lChild: &exprTreeNode{val: operators["*"]},
			rChild: &exprTreeNode{val: number(20)},
		},
		rChild: &exprTreeNode{val: number(500)},
	}, nil
}
