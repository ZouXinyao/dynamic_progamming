package bags

/* 动态规划+回溯，这种是不完全的动态规划 */
func completeBag01(W int, N int, w []int, v []int) int {
	dp := make([][]int, N + 1)
	for i:=0; i<N+1; i++ {
		dp[i] = make([]int, W + 1)
	}
	for i:=1; i<N+1; i++ {
		for j:=1; j<W+1; j++ {
			weight := w[i - 1]
			value  := v[i - 1]
			dp[i][j] = dp[i - 1][j]
			for k:=0; k<=j/weight; k++ {
				dp[i][j] = max(dp[i][j], dp[i - 1][j - weight * k] + value * k)
			}
		}
	}
	return dp[N][W]
}

/* 动态规划 */
func completeBag02(W int, N int, w []int, v []int) int {
	dp := make([][]int, N + 1)
	for i:=0; i<N+1; i++ {
		dp[i] = make([]int, W + 1)
	}
	for i:=1; i<N+1; i++ {
		for j:=1; j<W+1; j++ {
			weight := w[i - 1]
			value  := v[i - 1]
			if weight > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j - weight] + value)
			}
		}
	}
	return dp[N][W]
}

/* 动态规划，空间优化，滚动数组 */
func completeBag03(W int, N int, w []int, v []int) int {
	dp := make([][]int, 2)
	for i:=0; i<2; i++ {
		dp[i] = make([]int, W + 1)
	}
	for i:=1; i<N+1; i++ {
		for j:=1; j<W+1; j++ {
			idx := i % 2
			lastIdx := 1-idx
			weight := w[i - 1]
			value  := v[i - 1]
			if weight > j {
				dp[idx][j] = dp[lastIdx][j]
			} else {
				dp[idx][j] = max(dp[lastIdx][j], dp[idx][j - weight] + value)
			}
		}
	}
	return dp[N%2][W]
}
