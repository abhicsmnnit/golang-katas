package main

import (
	"errors"
	"fmt"
)

func main() {
	s := fmt.Sprintf("Hello") // highlighted by staticcheck tool - go vet lets it pass
	fmt.Println(s)

	err := returnErr(false)
	if err != nil {
		fmt.Println(err)
	}
	err = returnErr(true) // 2 issues highligted by staticcheck here
	fmt.Println("end of program")
}

func returnErr(hitMe bool) error {
	if hitMe {
		return errors.New("error has occurred")
	}
	return nil
}
