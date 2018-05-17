package lintcode

/**
 * Brute Force
 * @param k: An integer
 * @param n: An integer
 * @return: An integer denote the count of digit k in 1..n
 */
func digitCounts(k int, n int) int {
	counter := 0
	if k == 0 {
		counter++
	}
	for i := 1; i <= n; i++ {
		tmp := i
		for tmp != 0 {
			if tmp%10 == k {
				counter++
			}
			tmp = tmp / 10
		}
	}
	return counter
}
