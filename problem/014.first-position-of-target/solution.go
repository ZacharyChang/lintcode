package lintcode

/**
 * @param nums: The integer array.
 * @param target: Target to find.
 * @return: The first position of target. Position starts from 0.
 */
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for end-start > 1 {
		mid := (start + end) / 2
		if nums[mid] == target {
			end--
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
