package struct_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStruct(t *testing.T) {
	m := MyStruct[float32]{}

	assert.Equal(t, m.IsZero(), false)
}
