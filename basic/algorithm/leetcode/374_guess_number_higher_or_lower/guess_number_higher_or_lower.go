package guess_number_higher_or_lower

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		mid := left + (right-left)/2
		guessRes := guess(mid)
		if guessRes == 0 {
			return mid
		} else if guessRes == 1 {
			return dfs(mid+1, right)
		} else {
			return dfs(left, mid-1)
		}
	}
	return dfs(0, n)
}

func guess(num int) int {
	return 0
}
