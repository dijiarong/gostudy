package open_lock

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {
	fmt.Println(openLock([]string{"8888"}, "0009"))
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}
