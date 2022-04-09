package skiplist

type element struct {
	key   float64
	Value any
}

func (e *element) Update(val any) {
	e.Value = val
}

func newElement(key float64, val any) *element {
	return &element{
		key:   key,
		Value: val,
	}
}
