package map_func

// keys return the key of a map
// here m is generic using K and V
// V is contraint using any
// K is restrained using comparable i.e any type that supports != and == operation
func keys[K comparable, V any](m map[K]V) []K {
	key := make([]K, len(m))
	i := 0
	for k, _ := range m {
		key[i] = k
		i++
	}
	return key
}

// create a copy of map
// K and V are two different type with different contraints
func copy[K comparable, V any](m map[K]V) (copy map[K]V) {
	copy = map[K]V{}
	for k, v := range m {
		copy[k] = v
	}
	return
}

type Enteries[K, V any] struct {
	Key   K
	Value V
}

// Here a nested type parameter is used
// Enteries[K,V] intialize a new type and used here as return type
func enteries[K comparable, V any](m map[K]V) []*Enteries[K, V] {
	e := make([]*Enteries[K, V], len(m))
	i := 0
	for k, v := range m {
		newEntery := new(Enteries[K, V])
		newEntery.Key = k
		newEntery.Value = v
		e[i] = newEntery
		i++
	}
	return e
}
