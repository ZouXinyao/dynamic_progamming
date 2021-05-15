package solutions

// 62
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m + 1)
	for i := 0; i < m + 1; i++ {
		dp[i] = make([]int, n + 1)
	}
	dp[0][1] = 1
	for i := 1; i < m + 1; i++ {
		for j := 1; j < n + 1; j++ {
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}
	return dp[m][n]
}

// 63
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m + 1)
	for i := 0; i < m + 1; i++ {
		dp[i] = make([]int, n + 1)
	}
	dp[0][1] = 1
	for i := 1; i < m + 1; i++ {
		for j := 1; j < n + 1; j++ {
			if obstacleGrid[i - 1][j - 1] == 0 {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}
	return dp[m][n]
}

// 96
func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n + 1; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i - j - 1]
		}
	}
	return dp[n]
}