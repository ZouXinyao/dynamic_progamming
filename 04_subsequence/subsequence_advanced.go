package subsequence

func maxLongestRiseSubque(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	ans := 1
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			ans = max(ans, dp[i])
		}
	}
	return ans
}

func maxLongestRiseSubqueLength(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	ans := 1
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			ans = max(ans, dp[i])
		}
	}
	return ans
}
