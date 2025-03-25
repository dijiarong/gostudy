package my_calendar1

type MyCalendar struct {
	Tree *tree
}

func Constructor() MyCalendar {
	return MyCalendar{
		NewTree(0, 10^9),
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	// 读取
	if this.Tree.Query(startTime, endTime) != 0 {
		return false
	}
	this.Tree.Update(startTime, endTime, 1)
	return true
}

type tree struct {
	root *Node
}

type Node struct {
	left, right, val int
	Left, Right      *Node
	HasLazyAssign    bool
	assign           int
}

func NewTree(l, r int) *tree {
	return &tree{
		&Node{
			l, r, 0, nil, nil, false, 0,
		},
	}
}

func (t *tree) initChildNode(node *Node) {
	if node.left == node.right {
		// 叶子节点，无需划分
		return
	}
	if node.Left == nil || node.Right == nil {
		mid := node.left + (node.right-node.left)/2
		if node.Left == nil {
			node.Left = &Node{
				node.left, mid, 0, nil, nil, false, 0,
			}
		}
		if node.Right == nil {
			node.Right = &Node{
				mid + 1, node.right, 0, nil, nil, false, 0,
			}
		}
	}
}

func (t *tree) pushDown(root *Node) {
	if !root.HasLazyAssign {
		return
	}

	root.Left.assign = root.assign
	root.Left.HasLazyAssign = true
	root.Left.val = (root.Left.right - root.Left.left + 1) * root.assign

	root.Right.assign = root.assign
	root.Right.HasLazyAssign = true
	root.Right.val = (root.Right.right - root.Right.left + 1) * root.assign

	// 清除当前节点的懒标记
	root.HasLazyAssign = false
}

func (t *tree) Update(l, r, val int) {
	t.update(t.root, l, r, val)
}

func (t *tree) update(root *Node, l, r, val int) {
	if root.left >= l && root.right <= r {
		root.assign = val
		root.HasLazyAssign = true
		root.val = (r - l + 1) * val
		return
	}
	t.initChildNode(root)
	t.pushDown(root)
	mid := root.left + (root.right-root.left)/2
	if r <= mid {
		t.update(root.Left, l, r, val)
	} else if l > mid {
		t.update(root.Right, l, r, val)
	} else {
		t.update(root.Left, l, mid, val)
		t.update(root.Right, mid+1, r, val)
	}
	root.val = root.Left.val + root.Right.val
}

func (t *tree) Query(l, r int) int {
	return t.query(t.root, l, r)
}

func (t *tree) query(root *Node, l, r int) int {
	if root.left >= l && root.right <= r {
		return root.val
	}
	t.initChildNode(root)
	t.pushDown(root)
	mid := root.left + (root.right-root.left)/2
	if r <= mid {
		return t.query(root.Left, l, r)
	} else if l > mid {
		return t.query(root.Right, l, r)
	} else {
		return t.query(root.Left, l, mid) + t.query(root.Right, mid+1, r)
	}
}
