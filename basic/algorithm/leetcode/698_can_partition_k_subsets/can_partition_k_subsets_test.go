package can_partition_k_subsets

import (
	"fmt"
	"testing"
)

func TestCanPartitionKSubsets(t *testing.T) {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
