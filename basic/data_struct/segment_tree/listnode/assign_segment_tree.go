package segment_tree

type segmentNode struct {
	l, r          int
	Val           int
	Left, Right   *segmentNode
	LazyAssign    int
	HasLazyAssign bool
}

type AssignSegmentTree struct {
	Root       *segmentNode
	DefaultVal int
}

func NewAssignSegmentTree(start, end, defaultVal int) *AssignSegmentTree {
	return &AssignSegmentTree{
		Root: &segmentNode{
			l:   start,
			r:   end,
			Val: (end - start + 1) * defaultVal,
		},
		DefaultVal: defaultVal,
	}
}

// 动态创建子树
func (t *AssignSegmentTree) initChildNode(node *segmentNode) {
	if node.l == node.r {
		return
	}
	mid := node.l + (node.r-node.l)/2
	if node.Left == nil {
		node.Left = &segmentNode{
			l:   node.l,
			r:   mid,
			Val: t.DefaultVal,
		}
	}
	if node.Right == nil {
		node.Right = &segmentNode{
			r:   node.r,
			l:   mid + 1,
			Val: t.DefaultVal,
		}
	}
}

// lazy
func (t *AssignSegmentTree) pushDown(node *segmentNode) {
	if !node.HasLazyAssign {
		return
	}
	node.Left.HasLazyAssign = true
	node.Left.LazyAssign = node.LazyAssign
	node.Left.Val = (node.Left.r - node.Left.l + 1) * node.LazyAssign
	node.Right.HasLazyAssign = true
	node.Right.LazyAssign = node.LazyAssign
	node.Right.Val = (node.Right.r - node.Right.l + 1) * node.LazyAssign
	node.HasLazyAssign = false
}

func (t *AssignSegmentTree) Update(index, val int) {
	t.RangeUpdate(index, index, val)
}

func (t *AssignSegmentTree) RangeUpdate(start, end, val int) {
	t.rangeUpdate(t.Root, start, end, val)
}

func (t *AssignSegmentTree) rangeUpdate(node *segmentNode, start, end, val int) {
	if node.l >= start && node.r <= end {
		node.Val = (node.r - node.l + 1) * val
		node.HasLazyAssign = true
		node.LazyAssign = val
		return
	}
	t.initChildNode(node)
	t.pushDown(node)
	mid := node.l + (node.r-node.l)/2
	if end <= mid {
		t.rangeUpdate(node.Left, start, end, val)
	} else if start > mid {
		t.rangeUpdate(node.Right, start, end, val)
	} else {
		t.rangeUpdate(node.Left, start, mid, val)
		t.rangeUpdate(node.Right, mid+1, end, val)
	}
	node.Val = node.Left.Val + node.Right.Val
}

func (t *AssignSegmentTree) Query(l, r int) int {
	return t.query(t.Root, l, r)
}

func (t *AssignSegmentTree) query(node *segmentNode, l, r int) int {
	if node.l >= l && node.r <= r {
		return node.Val
	}
	t.initChildNode(node)
	t.pushDown(node)
	mid := node.l + (node.r-node.l)/2
	if r <= mid {
		return t.query(node.Left, l, r)
	} else if l > mid {
		return t.query(node.Right, l, r)
	} else {
		return t.query(node.Left, l, mid) + t.query(node.Right, mid+1, r)
	}
}
