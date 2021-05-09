package bags

func zeroOneBag(W int, N int, w []int, v []int) int {
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
				dp[i][j] = max(dp[i-1][j], dp[i - 1][j - weight] + value)
			}
		}
	}
	return dp[N][W]
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, i := range stones {
		sum += i
	}

	halfSum := sum / 2
	dp := make([][]int, len(stones) + 1)
	for i:=0; i<len(stones)+1; i++ {
		dp[i] = make([]int, halfSum + 1)
	}
	for i:=1; i<len(stones)+1; i++ {
		for j:=1; j<halfSum+1; j++ {
			weight := stones[i - 1]
			if weight > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i - 1][j - weight] + weight)
			}
		}
	}
	// 第一堆石头是sum - dp[len(stones)][halfSum]，第二堆石头是dp[len(stones)][halfSum]；
	return (sum - dp[len(stones)][halfSum]) - dp[len(stones)][halfSum]
}







