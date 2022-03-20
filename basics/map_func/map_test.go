package map_func

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	m := map[string]string{
		"first":  "w",
		"second": "t",
	}

	key := keys(m)

	sort.Strings(key)
	assert.Equal(t, []string{"first", "second"}, key)
}

func TestEnteries(t *testing.T) {
	m := map[string]string{
		"first":  "w",
		"second": "t",
	}

	key := enteries(m)

	for _, e := range key {
		assert.Equal(t, m[e.Key], e.Value)
	}

}

func TestCopy(t *testing.T) {
	m := map[string]string{
		"first":  "w",
		"second": "t",
	}
	assert.Equal(t, m, copy(m))
}
