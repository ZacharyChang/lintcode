package lintcode

/**
 * @param A: an array
 * @return: total of reverse pairs
 */
func reversePairs(A []int) int64 {
	counter := int64(0)
	if len(A) == 0 || len(A) == 1 {
		return 0
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				counter++
			}
		}
	}
	return counter
}
