package main

import (
	"errors"
	"fmt"
	"math"
)

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	diff1 := pair1.Val1.Diff(pair1.Val2)
	diff2 := pair2.Val1.Diff(pair2.Val2)

	if diff1 <= diff2 {
		return pair1
	} else {
		return pair2
	}
}

type Point2D struct {
	X, Y int
}

func (p Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p.X, p.Y)
}

func (p Point2D) Diff(from Point2D) float64 {
	diffX := p.X - from.X
	diffY := p.Y - from.Y

	return math.Sqrt(float64(diffX*diffX) + float64(diffY*diffY))
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p.X, p.Y, p.Z)
}

func (p Point3D) Diff(from Point3D) float64 {
	diffX := p.X - from.X
	diffY := p.Y - from.Y
	diffZ := p.Z - from.Z

	return math.Sqrt(float64(diffX*diffX) + float64(diffY*diffY) + float64(diffZ*diffZ))
}

// //////////////////////////

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr
}

func divAndRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}

	return num / denom, num % denom, nil
}

type Integer2 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr
	// ~ allows types which have their underlying types as mentioned above
}

func divAndRemainder2[T Integer2](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}

	return num / denom, num % denom, nil
}

func main() {
	pair1 := Pair[Point2D]{Point2D{10, 10}, Point2D{20, 20}}
	pair2 := Pair[Point2D]{Point2D{21, 21}, Point2D{30, 30}}
	fmt.Println(FindCloser(pair1, pair2))

	pair3 := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{20, 20, 20}}
	pair4 := Pair[Point3D]{Point3D{21, 21, 21}, Point3D{30, 30, 30}}
	fmt.Println(FindCloser(pair3, pair4))

	n1 := 10
	d1 := 3
	fmt.Println(divAndRemainder(n1, d1))

	n2 := 10_000_000_000_000_000
	d2 := 3_000_000_000_000_000
	fmt.Println(divAndRemainder(n2, d2))

	n3 := 10_000_000_000_000_000
	d3 := 0
	fmt.Println(divAndRemainder(n3, d3))

	type MyInt int
	mi1 := MyInt(10)
	mi2 := MyInt(3)
	// fmt.Println(divAndRemainder(mi1, mi2)) // compile error: MyInt does not satisfy Integer (possibly missing ~ for int in Integer)
	fmt.Println(divAndRemainder2(mi1, mi2))
}
