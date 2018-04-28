package lintcode

/**
 * @param nums: a rotated sorted array
 * @return: the minimum number in the array
 */
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	result := nums[0]
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] > nums[end] { // mid在右半边
			start = mid
		} else if nums[mid] < nums[end] { // mid在左半边
			end = mid
		} else {
			result = min(result, nums[mid]) // mid与end值相等时，位置因重复不确定
			end--
		}
	}
	// result = min(result, nums[start])
	result = min(result, nums[end])
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
