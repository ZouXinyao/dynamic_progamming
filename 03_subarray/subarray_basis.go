package subarray

import "math"

/* 647 */
func countSubstrings(s string) int {
	length := len(s)
	ans := 0
	dp := make([][]bool, length)
	for i:=0; i<length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
		ans++
	}

	for j:=1; j<length; j++ {
		for i:=0; i<j; i++ {
			dp[i][j] = false
			if s[i] == s[j] {
				if j - i <= 2 || dp[i + 1][j - 1] == true {
					dp[i][j] = true
					ans++
				}
			}
		}
	}
	return ans
}

/* 5 */
func longestPalindrome(s string) string {
	length := len(s)
	maxLen := 0
	idxs := [2]int{0,0}
	dp := make([][]int, length)
	for i:=0; i<length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = 1
	}

	for j:=1; j<length; j++ {
		for i:=0; i<j; i++ {
			if s[i] == s[j] {
				if j - i <= 2 {
					dp[i][j] = j - i + 1
				} else if dp[i + 1][j - 1] != 0 {
					dp[i][j] = dp[i + 1][j - 1] + 2
				}
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
				idxs[0] = i
				idxs[1] = j
			}
		}
	}
	return s[idxs[0]:idxs[1]+1]
}

/*剑指 Offer 42. 连续子数组的最大和*/
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	ans := nums[0]
	dp := make([]int, length)
	dp[0] = nums[0]
	for i:=1; i<length; i++ {
		dp[i] = math.MinInt32
		dp[i] = max(nums[i], nums[i] + dp[i - 1])
		ans = max(ans, dp[i])
	}
	return ans

}
