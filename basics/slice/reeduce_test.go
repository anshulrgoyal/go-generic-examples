package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	assert.Equal(t, 15, Reduce(v, func(acc int, ele int, i int, s []int) int {
		return acc + ele
	}))
}

func TestFilter(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}

	assert.Equal(t, []int{2, 4}, Filter(v, func(ele int, i int, s []int) bool {
		return ele%2 == 0
	}))
}

func TestMap(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	n := Map(v, func(a int, i int, arr []int) float32 {
		return float32(a * i)
	})
	assert.Equal(t, []float32{0, 2, 6, 12, 20}, n)
}
