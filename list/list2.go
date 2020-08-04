package list

type IntList2 struct {
	first *intListNode
	last  *intListNode
	len   int
}

func NewIntList2() *IntList2 {
	return &IntList2{nil, nil, 0}
}

func NewIntList2FromArray(a []int) *IntList2 {
	list := NewIntList2()
	for _, elem := range a {
		list.Add(elem)
	}
	return list
}

func (list *IntList2) Len() int {
	return list.len
}

func (list *IntList2) Add(elem int) {
	node := intListNode{nil, elem}
	if list.first == nil {
		list.first = &node
		list.last = &node
	} else {
		list.last.child = &node
		list.last = &node
	}
	list.len++
}

func (list *IntList2) Concat(other *IntList2) {
	if list.first == nil {
		*list = *other
	} else if other.Len() == 0 {
		// Do nothing
	} else {
		list.last.child = other.first
		list.last = other.last
		list.len += other.Len()
	}
}

func (list *IntList2) Each(f func(elem int)) {
	cur := list.first
	for cur != nil {
		f(cur.value)
		cur = cur.child
	}
}

func (list *IntList2) ToA() []int {
	a := make([]int, list.Len())
	{
		index := 0
		list.Each(func(elem int) {
			a[index] = elem
			index++
		})
	}
	return a
}
