package skiplist

import (
	"errors"
	"math/rand"
)

type Skiplist struct {
	layers []*node
}

func New() *Skiplist {
	sk := &Skiplist{}
	sk.layers = append(sk.layers, NewSentinal())
	return sk
}
func (sk *Skiplist) Insert(key float64, value any) {
	var table []*node
	layer := sk.layers[len(sk.layers)-1]
	for layer != nil {
		temp := layer.findInsertPos(key)
		table = append(table, temp)
		layer = temp.down
	}
	elem := newElement(key, value)
	lastNode := table[len(table)-1].insert(elem)
	for i := len(table) - 2; i >= 0; i-- {
		node := table[i]
		rand := rand.Float32()
		if rand >= 0.5 {
			nd := node.insert(elem)
			lastNode.up = nd
			nd.down = lastNode
			lastNode = nd
		} else {
			break
		}
	}
}

func (sk *Skiplist) Get(key float64) *element {
	node := sk.findNodeWithKey(key)
	if node != nil {
		return node.element
	}
	return nil
}
func (sk *Skiplist) Update(key float64, newValue any) error {
	node := sk.Get(key)
	if node == nil {
		return errors.New("key not found")
	}
	node.Update(newValue)
	return nil
}
func (sk *Skiplist) Delete(key float64) {}
func (sk *Skiplist) AddLayer() {
	lastLayer := sk.layers[len(sk.layers)-1]
	newLayer := NewSentinal()
	lastLayer.up = newLayer
	newLayer.down = lastLayer

	temp := lastLayer.next

	nlTemp := newLayer

	for temp != nil {
		rand := rand.Float32()
		if rand >= 0.5 {
			nlTemp = nlTemp.insert(temp.element)
			nlTemp.down = temp
			temp.up = nlTemp
		}
		temp = temp.next

	}
	sk.layers = append(sk.layers, newLayer)
}

func (sk *Skiplist) findNodeWithKey(key float64) *node {
	lastLayer := sk.layers[len(sk.layers)-1]
	for lastLayer != nil {
		temp := lastLayer.findInsertPos(key)
		if temp.key == key {
			lastLayer = temp
			break
		}
		lastLayer = temp.down
	}
	return lastLayer
}
