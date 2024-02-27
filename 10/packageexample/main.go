package main

import (
	"fmt"

	format "github.com/abhicsmnnit/golang-katas/10/packageexample/do-format"
	"github.com/abhicsmnnit/golang-katas/10/packageexample/math"
)

func main() {
	num := math.Triple(3)
	fmt.Println(format.Number(num))
}
