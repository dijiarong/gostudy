package segment_tree

type SegmentNode struct {
	l, r        int
	Val         int
	Left, Right *SegmentNode
}

type SegmentTree struct {
	root  *SegmentNode
	merge func(i, j int) int
}

func NewSegmentTree(arr []int, merge func(i, j int) int) *SegmentTree {
	return &SegmentTree{
		root:  build(arr, 0, len(arr)-1, merge),
		merge: merge,
	}
}

func build(arr []int, l, r int, merge func(i, j int) int) *SegmentNode {
	if l == r {
		return &SegmentNode{
			l:   l,
			r:   r,
			Val: arr[l],
		}
	}
	// 判断范围
	mid := l + (r-l)/2
	left := build(arr, l, mid, merge)
	right := build(arr, mid+1, r, merge)
	return &SegmentNode{
		l:     l,
		r:     r,
		Val:   merge(left.Val, right.Val),
		Left:  left,
		Right: right,
	}
}

func (t *SegmentTree) Query(l, r int) int {
	return t.query(t.root, l, r)
}

func (t *SegmentTree) query(node *SegmentNode, l, r int) int {
	if node.l == l && node.r == r {
		return node.Val
	}
	if node.l == node.r {
		return node.Val
	}
	mid := node.l + (node.r-node.l)/2
	if r <= mid {
		return t.query(node.Left, l, r)
	} else if l > mid {
		return t.query(node.Right, l, r)
	}
	return t.merge(
		t.query(node.Left, l, mid),
		t.query(node.Right, mid+1, r),
	)
}

func (t *SegmentTree) Update(index, val int) {
	t.update(t.root, index, val)
}

func (t *SegmentTree) update(node *SegmentNode, index, val int) {
	if node.l == node.r {
		node.Val = val
		return
	}
	mid := node.l + (node.r-node.l)/2
	if index <= mid {
		t.update(node.Left, index, val)
	} else if index > mid {
		t.update(node.Right, index, val)
	}
	node.Val = t.merge(node.Left.Val, node.Right.Val)
}
