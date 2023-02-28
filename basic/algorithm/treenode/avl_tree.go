package treenode

import "gostudy/basic/mymath"

// 任意节点的左子树高度与右子树高度差不超过1的二叉树
func IsAVLTree(root *TreeNode) bool {
	// 开始递归，从根节点开始
	res, _ := IsAVLTreeRec(root)
	return res
}

// 递归函数
// 参数：当前节点
// 返回值：bool 该节点是否为平衡二叉树 int 该节点的高度
// TODO: 可以优化，当某个节点不平衡时，结果已结束
func IsAVLTreeRec(root *TreeNode) (bool, int) {
	// 递归结束条件
	if root == nil {
		return true, 0
	}
	// 递归获取左子树的信息
	lB, lH := IsAVLTreeRec(root.Left)
	// 递归获取右子树的信息
	rB, rH := IsAVLTreeRec(root.Right)
	// 当前节点处理
	return lB && rB && lH-rH <= 1 && lH-rH >= -1, mymath.Math.GetMaxInt(lH, rH) + 1
}
