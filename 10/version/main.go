package main

import (
	"fmt"
	"log"
	"os"

	"github.com/learning-go-book-2e/simpletax"
	"github.com/shopspring/decimal"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("amount and zip must be provided")
	}

	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	zip := os.Args[2]
	percent, err := simpletax.TaxForZip(zip)
	if err != nil {
		log.Fatal(err)
	}

	total := amount.Add(amount.Mul(percent)).Round(2)
	fmt.Println(total)
}
