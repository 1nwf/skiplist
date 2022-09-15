package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var skiplist = New()

func TestInsert(t *testing.T) {
	for i := 0; i < 100; i++ {
		elem := skiplist.Insert(float64(i), i)
		assert.Equal(t, newElement(float64(i), i), elem)
	}
}

func TestGet(t *testing.T) {
	for i := 0; i < 100; i++ {
		elem := skiplist.Get(float64(i))
		assert.Equal(t, newElement(float64(i), i), elem)
	}
}

func TestAddLayer(t *testing.T) {
	for i := 1; i < 4; i++ {
		skiplist.AddLayer()
		assert.Equal(t, i+1, skiplist.Len())
	}
}
func TestDeleteLayer(t *testing.T) {
	for i := 1; i < 4; i++ {
		skiplist.RemoveLayer(0)
		assert.Equal(t, 4-i, skiplist.Len())
	}
}
