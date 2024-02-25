package main

import "fmt"

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

func Print[T Printable](t T) {
	fmt.Println(t.String())
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprintf("%d", mi)
}

type MyFloat64 float64

func (mf MyFloat64) String() string {
	return fmt.Sprintf("%f", mf)
}

type MyFloat32 float32

func (mf MyFloat32) String() string {
	return fmt.Sprintf("%f", mf)
}

func main() {
	Print(MyInt(10))
	Print(MyFloat64(64.64))
	// Print(MyFloat32(64.64)) // compile error: MyFloat32 does not satisfy Printable (MyFloat32 missing in ~int | ~float64)
}
