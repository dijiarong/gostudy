package successful_pairs_of_spells_and_potions

import (
	"fmt"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {
	spells := []int{1, 2, 3, 4, 5, 6, 7}
	potions := []int{1, 2, 3, 4, 5, 6, 7}
	var success int64 = 25
	fmt.Println(successfulPairs(spells, potions, success))
}
