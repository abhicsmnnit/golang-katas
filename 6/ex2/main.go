package main

import "fmt"

func main() {
	slice1 := []string{"hello", "hi"}
	UpdateSlice(slice1, "world")
	fmt.Println(slice1)

	slice2 := []string{"hello", "hi"}
	GrowSlice(slice2, "world")
	fmt.Println(slice2)
}

func UpdateSlice(slice []string, s string) {
	slice[len(slice)-1] = s
	fmt.Println(slice)
}

func GrowSlice(slice []string, s string) {
	slice = append(slice, s)
	fmt.Println(slice)
}
