package listimpl

type IntList1 struct {
	root *intListNode
	last *intListNode
	len  int
}

func NewIntList1() *IntList1 {
	root := intListNode{nil, 0}
	return &IntList1{&root, &root, 0}
}

func NewIntList1FromArray(a []int) *IntList1 {
	list := NewIntList1()
	for _, elem := range a {
		list.Add(elem)
	}
	return list
}

func (list *IntList1) Len() int {
	return list.len
}

func (list *IntList1) Add(elem int) {
	node := intListNode{nil, elem}
	list.last.child = &node
	list.last = &node
	list.len++
}

func (list *IntList1) Concat(other *IntList1) {
	if other.Len() == 0 {
		return
	}
	list.last.child = other.root.child
	list.last = other.last
	list.len += other.len
}

func (list *IntList1) Each(f func(elem int)) {
	cur := list.root
	for cur.child != nil {
		cur = cur.child
		f(cur.value)
	}
}

func (list *IntList1) ToA() []int {
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
