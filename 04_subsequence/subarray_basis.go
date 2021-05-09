package subsequence

/* 516 */
func longestPalindromeSubseq(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i:=0; i<length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = 1
	}
	for j:=1; j<length; j++ {
		for i:=j-1; i>=0; i-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i + 1][j - 1] + 2
			} else {
				dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[0][length - 1]
}

/* 1143 */
func longestCommonSubsequence(text1 string, text2 string) int {
	length1 := len(text1)
	length2 := len(text2)
	if length1 == 0 || length2 == 0 {
		return 0
	}
	dp := make([][]int, length1 + 1)
	for i:=0; i<length1+1; i++ {
		dp[i] = make([]int, length2 + 1)
	}
	for i := 1; i < length1 + 1; i++ {
		for j := 1; j < length2 + 1; j++ {
			dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
			if text1[i - 1] == text2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}
		}
	}
	return dp[length1][length2]
}

