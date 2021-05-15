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

// 494
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	length := len(nums)
	for _, n := range nums {
		sum += n
	}
	if (sum + target) % 2 == 1 {
		return 0
	}
	count := (sum + target) / 2
	dp := make([]int, count + 1)
	// 0放入0容量的背包。所以dp[0] = 1
	dp[0] = 1

	for i := 0; i < length; i++ {
		// 因为使用滚动数据，所以从后往前遍历
		for j := count; j >= nums[i]; j-- {
			dp[j] += dp[j - nums[i]]
		}
	}
	return dp[count]
}

// 377
func combinationSum4(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, target + 1)
	dp[0] = 1
	for i := 1; i < target + 1; i++ {
		for j := 0; j < length; j++ {
			if i >= nums[j] {
				dp[i] += dp[i - nums[j]]
			}
		}
	}
	return dp[target]
}

// 416
func canPartition(nums []int) bool {
	sum := 0
	for _, m := range nums {
		sum += m
	}
	if sum % 2 == 1{ return false }
	target := sum / 2
	dp := make([]bool, target + 1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j - nums[i]]
		}
	}
	return dp[target]
}





