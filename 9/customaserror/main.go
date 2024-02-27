package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.codes)
}

func (me MyErr) Codes() []int {
	return me.codes
}

func AFunctionThatReturnsAnError() error {
	return MyErr{codes: []int{1, 2, 3}}
}

func main() {
	err := AFunctionThatReturnsAnError()

	var me MyErr
	if errors.As(err, &me) {
		fmt.Println(me.codes)
	}

	var coder interface {
		Codes() []int
	}
	if errors.As(err, &coder) {
		fmt.Println(coder.Codes())
	}
}
