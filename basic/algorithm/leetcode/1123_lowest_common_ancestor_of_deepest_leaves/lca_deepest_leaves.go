package lowestcommonancestorofdeepestleaves

// 1123.最深叶节点的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, res := f(root)
	return res
}

// 获取高度,以及节点
func f(node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}
	lh, ln := f(node.Left)
	rh, rn := f(node.Right)
	if lh > rh {
		return lh + 1, ln
	}
	if rh > lh {
		return rh + 1, rn
	}
	return lh + 1, node
}
