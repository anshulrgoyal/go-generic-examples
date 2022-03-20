package struct_func

type MyStruct[T any] struct {
	inner T
}

func (m *MyStruct[T]) Get() T {
	return m.inner
}

func (m *MyStruct[T]) Set(v T) {
	m.inner = v
}

func (m *MyStruct[T]) IsZero() bool {

	switch f := (interface{})(m.inner).(type) {
	case int:
		return f == 0
	}

	return false
}
