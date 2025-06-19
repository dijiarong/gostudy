package remove_nth_node_from_end_of_list

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	removeNthFromEnd(makeListNodes([]int{1}), 1)
	removeNthFromEnd(makeListNodes([]int{1, 2, 3, 4, 5}), 2)
	removeNthFromEnd(makeListNodes([]int{1, 2}), 1)
}
