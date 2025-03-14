package segment_tree

// 动态的创建左右子树
type DynamicSegmentTree struct {
	Root       *SegmentNode
	merge      func(i, j int) int
	defaultVal int
}

func NewDynamicSegmentTree(
	start, end, defaultVal int, merge func(i, j int) int) *DynamicSegmentTree {
	return &DynamicSegmentTree{
		Root: &SegmentNode{
			l:   start,
			r:   end,
			Val: defaultVal,
		},
		merge:      merge,
		defaultVal: defaultVal,
	}
}

func (t *DynamicSegmentTree) initChildNode(node *SegmentNode) {
	if node.l == node.r {
		return
	}
	mid := node.l + (node.r-node.l)/2
	if node.Left == nil {
		node.Left = &SegmentNode{
			l:   node.l,
			r:   mid,
			Val: t.defaultVal,
		}
	}
	if node.Right == nil {
		node.Right = &SegmentNode{
			l:   mid + 1,
			r:   node.r,
			Val: t.defaultVal,
		}
	}
}

func (t *DynamicSegmentTree) Update(index, val int) {
	t.update(t.Root, index, val)
}

func (t *DynamicSegmentTree) update(node *SegmentNode, index, val int) {
	if node.l == node.r {
		node.Val = val
		return
	}
	t.initChildNode(node)
	mid := node.l + (node.r-node.l)/2
	if index <= mid {
		t.update(node.Left, index, val)
	} else {
		t.update(node.Right, index, val)
	}
	node.Val = t.merge(node.Left.Val, node.Right.Val)
}

func (t *DynamicSegmentTree) Query(l, r int) int {
	return t.query(t.Root, l, r)
}

func (t *DynamicSegmentTree) query(node *SegmentNode, ql, qr int) int {
	if node.l >= ql && node.r <= qr {
		return node.Val
	}
	mid := node.l + (node.r-node.l)/2
	if qr <= mid {
		return t.query(node.Left, ql, mid)
	} else if ql > mid {
		return t.query(node.Right, mid+1, qr)
	} else {
		return t.merge(t.query(node.Left, ql, mid), t.query(node.Right, mid+1, qr))
	}
}
