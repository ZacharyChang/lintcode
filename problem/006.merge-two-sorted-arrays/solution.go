package lintcode

/**
 * solution 1: recursion
 * @param A: sorted integer array A
 * @param B: sorted integer array B
 * @return: A new sorted integer array
 */
func mergeSortedArray(A []int, B []int) []int {
	if A == nil || len(A) == 0 {
		return B
	}
	if B == nil || len(B) == 0 {
		return A
	}
	if len(A) < len(B) {
		A, B = B, A
	}
	if A[0] >= B[0] {
		nextArray := mergeSortedArray(A, B[1:])
		return append([]int{B[0]}, nextArray...)
	} else {
		nextArray := mergeSortedArray(A[1:], B)
		return append([]int{A[0]}, nextArray...)
	}
}

/**
 * solution 2: two pointers
 * @param A: sorted integer array A
 * @param B: sorted integer array B
 * @return: A new sorted integer array
 */
func mergeSortedArray_2(A []int, B []int) []int {
	result := make([]int, len(A)+len(B))
	i, j, k := 0, 0, 0
	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			result[k] = A[i]
			i++
		} else {
			result[k] = B[j]
			j++
		}
		k++
	}
	if i < len(A) {
		copy(result[k:], A[i:])
	}
	if j < len(B) {
		copy(result[k:], B[j:])
	}
	return result
}
