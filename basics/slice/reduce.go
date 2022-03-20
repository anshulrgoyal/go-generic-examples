package slice

func Reduce[T any, R any](s []T, f func(acc R, ele T, i int, s []T) R) R {
	var acc R
	for i, value := range s {
		acc = f(acc, value, i, s)
	}
	return acc
}

func Filter[T any](s []T, f func(ele T, i int, s []T) bool) []T {
	var mappedArray []T
	for i, value := range s {
		if f(value, i, s) {
			mappedArray = append(mappedArray, value)
		}

	}
	return mappedArray
}

// Map return a copy of array with an transomation performed
// it take a func which perform transformatio
// here there are 2 type parameters T the type of Slice given to map and R the type of value returned by transformation function
func Map[T any, R any](s []T, f func(a T, i int, arr []T) R) []R {
	mappedArray := make([]R, len(s))
	for i, value := range s {
		mappedArray[i] = f(value, i, s)
	}
	return mappedArray
}

func Every[T any](s []T, f func(ele T, i int, arr []T) bool) bool {
	r := true
	for i, value := range s {
		r = r && f(value, i, s)
	}
	return r
}

func Some[T any](s []T, f func(ele T, i int, arr []T) bool) bool {
	r := false
	for i, value := range s {
		r = r || f(value, i, s)
	}
	return r
}


func Flat[T any](s [][]T) []T {
	l:=0
	for _,ele:= range s {
		l =l + len(ele)
	}
	arr:=make([]T,0,l)
	for _,ele:= range s {
		for _,e:=range ele {
			arr=append(arr,e)
		}
	}
	return arr
}