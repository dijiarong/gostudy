package searchsuggestionssystem

import (
	"fmt"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWord := "mouse"
	fmt.Print(suggestedProducts(products, searchWord))
	fmt.Print("\n")
}
