package lintcode

/**
 * @param matrix: A list of lists of integers
 * @param target: An integer you want to search in matrix
 * @return: An integer indicate the total occurrence of target in the given matrix
 */
func searchMatrix(matrix [][]int, target int) int {
	if len(matrix) == 0 {
		return 0
	}

	count := 0
	start := 0
	end := len(matrix) - 1
	row := end

	for start < end-1 {
		mid := (start + end) / 2
		if matrix[mid][0] == target {
			row = mid
			break
		} else if matrix[mid][0] > target {
			end = mid
			row = mid
		} else if matrix[mid][0] < target {
			start = mid
		}
	}
	for i := 0; i <= row; i++ {
		if binarySearch(matrix[i], target) {
			count++
		}
	}
	return count
}

func binarySearch(array []int, target int) bool {
	if len(array) == 0 {
		return false
	}
	if array[0] > target || array[len(array)-1] < target {
		return false
	}
	if array[0] == target || array[len(array)-1] == target {
		return true
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
	return false
}
