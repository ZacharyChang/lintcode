package lintcode

/**
 * @param n: an integer
 * @return: an ineger f(n)
 */
func fibonacci(n int) int {
	array := make([]int, n+2)
	array[0] = 0
	array[1] = 1
	for i := 2; i < n; i++ {
		array[i] = array[i-2] + array[i-1]
	}
	return array[n-1]
}
