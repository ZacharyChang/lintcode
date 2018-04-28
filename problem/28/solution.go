package lintcode

/**
 * @param matrix: matrix, a list of lists of integers
 * @param target: An integer
 * @return: a boolean, indicate whether matrix contains target
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	start := 0
	end := len(matrix) - 1
	for start < end-1 {
		mid := (start + end) / 2
		if matrix[mid][0] > target {
			end = mid
		} else if matrix[mid][0] < target {
			start = mid
		} else {
			return true
		}
	}
	return binarySearch(matrix[start], target) || binarySearch(matrix[end], target)
}

func binarySearch(array []int, target int) bool {
	if len(array) == 0 {
		return false
	}
	start := 0
	end := len(array) - 1
	for start < end-1 {
		mid := (start + end) / 2
		if array[mid] == target {
			return true
		} else if array[mid] > target {
			end = mid
		} else if array[mid] < target {
			start = mid
		}
	}
	return array[start] == target || array[end] == target
}
