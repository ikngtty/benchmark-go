package anylist

type anyListNode struct {
	child *anyListNode
	value interface{}
}

type AnyList struct {
	root *anyListNode
	last *anyListNode
	len  int
}

func NewAnyList() *AnyList {
	root := anyListNode{nil, nil}
	return &AnyList{&root, &root, 0}
}

func (list *AnyList) Len() int {
	return list.len
}

func (list *AnyList) Add(elem interface{}) {
	node := anyListNode{nil, elem}
	list.last.child = &node
	list.last = &node
	list.len++
}

func (list *AnyList) Each(f func(elem interface{})) {
	cur := list.root
	for cur.child != nil {
		cur = cur.child
		f(cur.value)
	}
}
