package main

import (
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

func main() {
	pair1 := Pair[Point2D]{Point2D{10, 10}, Point2D{20, 20}}
	pair2 := Pair[Point2D]{Point2D{21, 21}, Point2D{30, 30}}
	fmt.Println(FindCloser(pair1, pair2))

	pair3 := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{20, 20, 20}}
	pair4 := Pair[Point3D]{Point3D{21, 21, 21}, Point3D{30, 30, 30}}
	fmt.Println(FindCloser(pair3, pair4))
}
