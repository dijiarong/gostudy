package segment_tree

type SegmentTree struct {
	arr   []int
	count int
	merge func(i, j int) int
}

func NewArraySegmentTree(arr []int, merge func(i, j int) int) *SegmentTree {
	tree := &SegmentTree{
		arr:   make([]int, len(arr)*4),
		count: len(arr),
		merge: merge,
	}
	tree.build(arr, 0, 0, len(arr)-1)
	return tree
}

func (t *SegmentTree) build(arr []int, rootIndex, l, r int) {
	if l == r {
		t.arr[rootIndex] = arr[l]
		return
	}
	leftChildRootIndex := t.leftChildRootIndex(rootIndex)
	rightChildRootIndex := t.rightChildRootIndex(rootIndex)
	mid := l + (r-l)/2
	t.build(arr, leftChildRootIndex, l, mid)
	t.build(arr, rightChildRootIndex, mid+1, r)
	t.arr[rootIndex] = t.merge(t.arr[leftChildRootIndex], t.arr[rightChildRootIndex])
}

func (t *SegmentTree) Update(index, val int) {
	t.update(0, 0, t.count-1, index, val)
}

func (t *SegmentTree) update(rootIndex, l, r, index, val int) {
	if l == r {
		t.arr[rootIndex] = val
		return
	}
	mid := l + (r-l)/2
	if index <= mid {
		t.update(t.leftChildRootIndex(rootIndex), l, mid, index, val)
	} else if index > l {
		t.update(t.rightChildRootIndex(rootIndex), mid+1, r, index, val)
	}
	t.arr[rootIndex] = t.merge(t.arr[t.leftChildRootIndex(rootIndex)], t.arr[t.rightChildRootIndex(rootIndex)])
}

func (t *SegmentTree) Query(l, r int) int {
	return t.query(0, 0, t.count-1, l, r)
}

func (t *SegmentTree) query(rootIndex, l, r, ql, qr int) int {
	if l == r {
		return t.arr[rootIndex]
	}
	mid := l + (r-l)/2
	leftChildRootIndex := t.leftChildRootIndex(rootIndex)
	rightChildRootIndex := t.rightChildRootIndex(rootIndex)
	if qr <= mid {
		return t.query(leftChildRootIndex, l, mid, ql, qr)
	} else if ql > mid {
		return t.query(rightChildRootIndex, mid+1, r, ql, qr)
	} else {
		return t.merge(t.query(leftChildRootIndex, l, mid, ql, mid), t.query(rightChildRootIndex, mid+1, r, mid+1, qr))
	}
}

// 左子树的根索引
func (t *SegmentTree) leftChildRootIndex(i int) int {
	return 2*i + 1
}

// 右子树的根索引
func (t *SegmentTree) rightChildRootIndex(i int) int {
	return 2*i + 2
}
