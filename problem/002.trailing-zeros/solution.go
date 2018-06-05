package lintcode

var array []int

/**
 * @param n: A long integer
 * @return: An integer, denote the number of trailing zeros in n!
 */
func trailingZeros(n int64) int64 {
	if n == 0 {
		return 0
	}
	var count int64
	for i := n; i > 0; {
		i = i / 5
		count += int64(i)
	}
	// // 因数5的个数必大于因数2的个数
	// for m%5 == 0 && m > 0 {
	// 	count++
	// 	m = m / 5
	// }
	// for i := 1; i < n; i++ {

	// 	array[i] = array[i-1]
	// }
	return count
}
