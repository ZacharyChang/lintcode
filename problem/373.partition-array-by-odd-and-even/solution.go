package lintcode

func partitionArray(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		for i < j && nums[i]%2 == 1 {
			i++
		}
		for i < j && nums[j]%2 == 0 {
			j--
		}
		if i < j {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
		}
	}
	return nums
}
