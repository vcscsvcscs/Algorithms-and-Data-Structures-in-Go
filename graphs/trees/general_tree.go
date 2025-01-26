package trees

type Node struct {
	Key     any
	Sibling *Node
	Child   *Node
}

func NewNode(key any) *Node {
	return &Node{Key: key}
}

func (t *Node) PreOrder(process func(*Node)) {
	for t != nil {
		process(t)
		t.Child.PreOrder(process)
		t = t.Sibling
	}
}

func (t *Node) PostOrder(process func(*Node)) {
	for t != nil {
		t.Child.PostOrder(process)
		process(t)
		t = t.Sibling
	}
}
