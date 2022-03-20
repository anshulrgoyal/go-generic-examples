package variable

import (
	"fmt"
	"testing"
)

func simpleFunc[T any](a T) T {
	return a
}

// intialize as a global vairable in package
var simpleFuncFloat = simpleFunc[float32]

func TestIntialized(t *testing.T) {
	// generic function can be intialized with a type and used as varaible
	b := simpleFunc[int]

	// calling the function
	fmt.Println(b(1))

	fmt.Println(simpleFuncFloat(3.0))
}
