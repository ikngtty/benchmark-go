package anylist

type intListNode struct {
	child *intListNode
	value int
}

type IntList struct {
	root *intListNode
	last *intListNode
	len  int
}

func NewIntList() *IntList {
	root := intListNode{nil, 0}
	return &IntList{&root, &root, 0}
}

func (list *IntList) Len() int {
	return list.len
}

func (list *IntList) Add(elem int) {
	node := intListNode{nil, elem}
	list.last.child = &node
	list.last = &node
	list.len++
}

func (list *IntList) Each(f func(elem int)) {
	cur := list.root
	for cur.child != nil {
		cur = cur.child
		f(cur.value)
	}
}
