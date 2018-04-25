package lintcode

/**
 * 二分查找法
 * @param nums: a rotated sorted array
 * @return: the minimum number in the array
 */
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return min(nums[0], nums[1])
	}
	start := 0
	end := len(nums) - 1
	for start < end-1 {
		// 如果最后一个元素大于第一个元素，说明有序，最小元素即为第一个
		if nums[start] < nums[end] {
			return nums[start]
		}
		middle := (start + end) / 2
		if nums[middle] < nums[start] {
			end = middle
		} else if nums[middle] > nums[start] {
			start = middle
		}
	}
	return min(nums[start], nums[end])
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
