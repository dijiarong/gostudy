package hard

import (
	"gostudy/basic/algorithm/listnode"
	"sort"
)

// 合并k个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/
func MergeKLists(lists []*listnode.ListNode) *listnode.ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	// 维护一个最小堆
	heapTmp := make([]*listnode.ListNode, 0, len(lists))
	for _, v := range lists {
		if v != nil {
			heapTmp = append(heapTmp, v)
		}
	}
	// 将头节点放入堆中
	sort.Slice(heapTmp, func(i, j int) bool {
		return heapTmp[i].Val < heapTmp[j].Val
	})
	head := &listnode.ListNode{Val: -1, Next: nil}
	curNode := head
	for len(heapTmp) != 0 {
		curNode.Next = heapTmp[0]
		curNode = curNode.Next
		heapTmp = heapTmp[1:]
		if curNode.Next != nil {
			heapTmp = append(heapTmp, curNode.Next)
		}
		sort.Slice(heapTmp, func(i, j int) bool {
			return heapTmp[i].Val < heapTmp[j].Val
		})
	}
	return head.Next
}
