package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ListNode struct {
	Key         byte
	Val         int
	left, right *ListNode
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Fields(str)
	aft, mid := strs[0], strs[1]
	midMap := map[byte]int{}
	for i := range mid {
		midMap[mid[i]] = i
	}
	var makeTree func(l1, r1, l2, r2 int) *ListNode
	makeTree = func(l1, r1, l2, r2 int) *ListNode {
		// 结束条件
		if l1 > r1 || l2 > r2 {
			return nil
		}
		// r1为后序遍历的最后一个节点，也就是当地树的根节点
		cur := &ListNode{
			Key: aft[r1],
		}
		if l1 == r1 {
			return cur
		}
		// 查找在中序遍历里的位置，从而缺点左右子树的范围
		midIndex := midMap[aft[r1]]
		// 递归拼接左子树
		cur.left = makeTree(l1, l1+midIndex-1-l2, l2, midIndex-1)
		// 递归拼接右子树
		cur.right = makeTree(l1+midIndex-l2, r1-1, midIndex+1, r2)
		return cur
	}
	node := makeTree(0, len(aft)-1, 0, len(mid)-1)
	stack := []*ListNode{node}
	b := strings.Builder{}
	for len(stack) != 0 {
		num := len(stack)
		for i := 0; i < num; i++ {
			cur := stack[i]
			b.WriteByte(cur.Key)
			if cur.left != nil {
				stack = append(stack, cur.left)
			}
			if cur.right != nil {
				stack = append(stack, cur.right)
			}
		}
		stack = stack[num:]
	}
	fmt.Println(b.String())
}
