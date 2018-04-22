package lintcode

/**
 * @param n: An integer
 * @return: An integer
 */
func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	array := make([]int, n)
	array[0] = 1
	array[1] = 2
	for i := 2; i < n; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array[n-1]
}
