package max_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root TreeNode) int {
	res := 0
	return maxPathSum1(&root, &res)
}

// 递归返回包含当前节点的左 右路径的最大值
func maxPathSum1(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := max(maxPathSum1(root.Left, res), 0)
	r := max(maxPathSum1(root.Right, res), 0)
	cur := l + r + root.Val
	*res = max(*res, cur)
	return max(l, r) + root.Val
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
