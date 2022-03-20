package main

import (
	"constraints"
	"fmt"
)

type comparable interface {
	~int | ~float32 | ~int8 | ~int16 | ~int32 | ~int64 | ~float64 | ~string
}

type Size int32

type Point[T comparable, F any] struct {
	r F
	X T
	Y T
}

func (point *Point[T, F]) min(x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}

func min[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	{
		var x, y Size
		x = 0
		y = 1
		println(min(x, y))
	}
	{
		var x, y float32
		x = 0
		y = 1
		println(min(x, y))
	}
	{
		var x, y string
		x = "a"
		y = "b"
		println(min(x, y))
	}
	TestIntialized()
}

func simpleFunc[T any](a T) T {
	return a
}

// intialize as a global vairable in package
var simpleFuncFloat = simpleFunc[float32]

func TestIntialized() {
	// generic function can be intialized with a type and used as varaible
	b := simpleFunc[int]

	// calling the function
	fmt.Println(b(1))

	fmt.Println(simpleFuncFloat(3.0))
}
