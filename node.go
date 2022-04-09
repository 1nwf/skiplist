package skiplist

import (
	"math"
)

type node struct {
	*element
	next *node
	up   *node
	prev *node
	down *node
}

func newNode(e *element) *node {
	return &node{
		element: e,
	}
}
func (n *node) insert(e *element) *node {
	nd := newNode(e)
	temp := n.findInsertPos(e.key)
	nd.next = temp.next
	nd.prev = temp
	temp.next = nd
	return nd
}

func (n *node) findInsertPos(key float64) *node {
	temp := n
	for temp.next != nil && temp.key < key {
		temp = temp.next
	}
	if key >= temp.key {
		return temp
	}
	return temp.prev
}

func NewSentinal() *node {
	return &node{
		element: &element{
			key: math.Inf(-1),
		},
	}
}
