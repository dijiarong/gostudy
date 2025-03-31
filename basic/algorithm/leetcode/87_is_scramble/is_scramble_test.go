package is_scramble

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	s1 := "great"
	s2 := "rgeat"
	fmt.Println(isScramble(s1, s2))
}
