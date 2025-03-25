package range_sum_query_mutable

type NumArray struct {
	Tree *SegmentTree
}

func Constructor(nums []int) NumArray {
	t := NewSegmentTree(0, len(nums)-1, 0)
	for i, num := range nums {
		t.Update(i, num)
	}
	return NumArray{
		Tree: t,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.Tree.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Tree.Query(left, right)
}

type SegmentNode struct {
	l, r        int
	Val         int
	Left, Right *SegmentNode
}

type SegmentTree struct {
	Root       *SegmentNode
	DefaultVal int
}

func NewSegmentTree(
	l, r, defaultVal int,
) *SegmentTree {
	return &SegmentTree{
		Root: &SegmentNode{
			l, r, (r - l + 1) * defaultVal, nil, nil,
		},
		DefaultVal: defaultVal,
	}
}

func (t *SegmentTree) initChildNode(root *SegmentNode) {
	if root.l == root.r {
		return
	}
	mid := root.l + (root.r-root.l)/2
	if root.Left == nil {
		root.Left = &SegmentNode{
			l:   root.l,
			r:   mid,
			Val: (mid - root.l + 1) * t.DefaultVal,
		}
	}
	if root.Right == nil {
		root.Right = &SegmentNode{
			l:   mid + 1,
			r:   root.r,
			Val: (root.r - mid) * t.DefaultVal,
		}
	}
}

func (t *SegmentTree) Update(index, val int) {
	t.update(t.Root, index, val)
}

func (t *SegmentTree) update(root *SegmentNode, index, val int) {
	if root.l == root.r {
		root.Val = val
		return
	}
	t.initChildNode(root)
	mid := root.l + (root.r-root.l)/2
	if index <= mid {
		t.update(root.Left, index, val)
	} else {
		t.update(root.Right, index, val)
	}
	root.Val = root.Left.Val + root.Right.Val
}

func (t *SegmentTree) Query(l, r int) int {
	return t.query(t.Root, l, r)
}

func (t *SegmentTree) query(root *SegmentNode, l, r int) int {
	if root.l == l && root.r == r {
		return root.Val
	}
	mid := root.l + (root.r-root.l)/2
	if r <= mid {
		return t.query(root.Left, l, r)
	} else if l > mid {
		return t.query(root.Right, l, r)
	} else {
		return t.query(root.Left, l, mid) + t.query(root.Right, mid+1, r)
	}
}
