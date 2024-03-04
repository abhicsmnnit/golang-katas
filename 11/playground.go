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

	true := false // can be highlighted by revive tool and certain config files. Run `revive -config built_in.toml ./...`.
	fmt.Println(true)
}

func returnErr(hitMe bool) error {
	if hitMe {
		return errors.New("error has occurred")
	}
	return nil
}
