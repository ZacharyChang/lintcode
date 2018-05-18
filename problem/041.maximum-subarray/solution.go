package lintcode

/**
 * @param nums: A list of integers
 * @return: A integer indicate the sum of max subarray
 */
func maxSubArray(nums []int) int {
	maxNum := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		// dp[i+1]=max(dp[i],dp[i]+dp[i+1])
		maxNum = max(maxNum+nums[i], nums[i])
		if maxNum > result {
			result = maxNum
		}
	}
	return result
}

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}
