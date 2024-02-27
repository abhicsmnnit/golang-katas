package main

import (
	"errors"
	"fmt"
)

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if target, ok := target.(ResourceErr); ok {
		ignoreResource := target.Resource == ""
		ignoreCode := target.Code == 0
		matchResource := target.Resource == re.Resource
		matchCode := target.Code == re.Code

		return matchResource && matchCode ||
			ignoreResource && matchCode ||
			matchResource && ignoreCode
	}

	return false
}

func main() {
	err1 := ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	err2 := ResourceErr{
		Resource: "Network",
		Code:     1000,
	}

	if errors.Is(err1, ResourceErr{Resource: "Database"}) {
		fmt.Println("Database error: ", err1)
	}

	if errors.Is(err2, ResourceErr{Resource: "Database"}) {
		fmt.Println("Database error: ", err2)
	}

	if errors.Is(err1, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err1)
	}

	if errors.Is(err2, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err2)
	}

	if errors.Is(err1, ResourceErr{Resource: "Database", Code: 123}) {
		fmt.Println("Database is broken and Code 123 triggered:", err1)
	}

	if errors.Is(err1, ResourceErr{Resource: "Network", Code: 123}) {
		fmt.Println("Network is broken and Code 123 triggered:", err1)
	}
}
