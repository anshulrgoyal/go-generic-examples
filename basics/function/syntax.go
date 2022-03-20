package function

// reverse function retruns a reversed copy of array
// T type parameter specify any type i.e it can be int, struct, slice, string etc.
// It return the same type of slice
// since we don't use any method on T so we can use any type here
// but if we used some operation here interface constraint needs to be changed
func reverse[T any](a []T) []T {

	// create a slice of T with length of slice a
	reversedArray := make([]T, len(a))

	// loop till middle of a and write to reversedArray
	for i := 0; i <= len(a)/2; i++ {
		reversedArray[i] = a[len(a)-i-1]
		reversedArray[len(a)-i-1] = a[i]
	}

	// return reversedArray
	return reversedArray
}

// stringer just define a method to convert a value to string
type stringer interface {
	String() string
}

// ToString convert a type to string
// it used S as stringer type that will constarint allowed in the function to always have a String method
// This is compile time type checked
func ToString[S stringer](s S) string {
	return s.String()
}

/*
here ToString will not work since no String method on any
func ToString[S any](s S) string{
	return s.String() // INVALID: compiler will throw error that String method doesn't exist on any
}

*/
