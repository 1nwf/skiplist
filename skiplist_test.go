package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sklist = New()

func TestInsert(t *testing.T) {
	for i := 0; i < 10; i++ {
		elem := sklist.Insert(float64(i), i)
		assert.Equal(t, newElement(float64(i), i), elem)
	}
}

func TestGet(t *testing.T) {
	for i := 0; i < 10; i++ {
		elem := sklist.Get(float64(i))
		assert.Equal(t, newElement(float64(i), i), elem)
	}
}

func TestAdLayer(t *testing.T) {
	for i := 1; i < 4; i++ {
		sklist.AddLayer()
		assert.Equal(t, i+1, sklist.Len())
	}
}
func TestDeleteLayer(t *testing.T) {
	for i := 1; i < 4; i++ {
		sklist.RemoveLayer(0)
		assert.Equal(t, 4-i, sklist.Len())
	}
}
