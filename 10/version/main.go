package main

import (
	"fmt"
	"log"
	"os"

	"github.com/learning-go-book-2e/simpletax/v2"
	"github.com/shopspring/decimal"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("amount, zip and country-code must be provided")
	}

	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	zip := os.Args[2]
	countryCode := os.Args[3]
	percent, err := simpletax.ForCountryPostalCode(countryCode, zip)
	if err != nil {
		log.Fatal(err)
	}

	total := amount.Add(amount.Mul(percent)).Round(2)
	fmt.Println(total)
}
