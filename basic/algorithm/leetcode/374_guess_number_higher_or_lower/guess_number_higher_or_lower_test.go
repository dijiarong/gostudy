package guess_number_higher_or_lower

import (
	"fmt"
	"testing"
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func TestGuessNumber(t *testing.T) {
	fmt.Println(guessNumber(2126753390))
}
