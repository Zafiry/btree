package btree

type internalNode struct {
	baseNode
	children []node
}

func newInternalNode(t *btree, p node, r uint64) *internalNode {
	return &internalNode{
		tree:     t,
		parent:   p,
		revision: r,
		keys:     make([][]byte, 0, t.order-1),
		children: make([]node, 0, t.order),
	}
}

func (n *internalNode) isLeaf() bool {
	return false
}

func (n *internalNode) getValue([]byte) []byte {
	exist, idx := n.findPos(key)
	// equal: go to right
	if exist {
		return n.children[idx+1].getValue(key)
	}
	return n.children[idx].getValue(key)
}

func (n *internalNode) insertChildAt(idx int, child node) {
	lastCnt := len(n.children)
	n.children = append(n.children, nil)
	// move backward
	for i := lastCnt - 1; i >= idx; i-- {
		n.children[i+1] = n.children[i]
	}
	n.children[idx] = child
}

func (n *internalNode) clone(revision uint64) *internalNode {
	newINode := newInternalNode(n.tree, n.parent, revision)
	for _, k := range n.keys {
		newINode.keys = append(newINode.keys, k)
	}
	for _, c := range n.children {
		newINode.children = append(newINode.children, c)
	}
}

func (n *internalNode) insert(key, value []byte, revision uint64) *insertResult {
	//TODO:
	return nil
}