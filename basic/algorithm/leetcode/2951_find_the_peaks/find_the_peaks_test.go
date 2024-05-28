package findthepeaks

import (
	"fmt"
	"testing"
)

func TestFindPeaks(t *testing.T) {
	res := []int{1, 4, 3, 8, 5}
	fmt.Print(findPeaks(res))
}
