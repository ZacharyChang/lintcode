package lintcode

/**
 * @param n: An integer
 * @param nums: An array
 * @return: the Kth largest element
 */
func kthLargestElement(n int, nums []int) int {
	quickSort(nums)
	return nums[len(nums)-n]
}

func quickSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		for j-1 >= 0 && j > i && nums[j] >= nums[0] {
			j--
		}
		for i+1 < len(nums) && j > i && nums[i] <= nums[0] {
			i++
		}
		if i == j {
			break
		} else {
			swap(nums, i, j)
		}

	}
	swap(nums, 0, j)
	quickSort(nums[0:j])
	quickSort(nums[j+1 : len(nums)])
}

func swap(nums []int, i int, j int) {
	if i == j {
		return
	}
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
