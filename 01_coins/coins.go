package coins

import (
	"math"
	"sort"
)

/* 贪心算法，这种方法只会达到局部最优，认为局部最优=全局最优 */
func getMinCoinCount01(total int, values []int) int {
	res := 0
	sort.Slice(
		values,
		func(i, j int) bool {
			return values[i] > values[j]
		},
	)
	for i := 0; i < len(values); i++ {
		currentCount := total / values[i]
		if currentCount == 0 {
			continue
		}
		res += currentCount
		total -= values[i] * currentCount
		if total == 0 {
			return res
		}
	}
	return -1
}

/* 递归枚举，这种方法枚举所有可能的情况，返回全局最优解 */
func getMinCoinCount02(total int, values []int) int {
	if total == 0 {
		return 0
	}

	minCount := math.MaxInt32
	for i := 0; i < len(values); i++ {
		if total < values[i] {
			continue
		}
		rest := total - values[i]
		restCount := getMinCoinCount02(rest, values)
		if restCount == -1 {
			continue
		}

		currentCount := restCount + 1
		if currentCount < minCount {
			minCount = currentCount
		}
	}
	if minCount == math.MinInt32 {
		return -1
	}
	return minCount
}

/* 递归枚举+备忘录，记录中间结果，避免重复计算 */
func getMinCoinCount03(total int, values []int) int {
	memo := make([]int, total + 1)
	for i := 0; i < total+1; i++ {
		memo[i] = -2
	}
	memo[0] = 0

	return helper03(total, values, memo)
}

func helper03(total int, values []int, memo []int) int {
	if memo[total] != -2 {
		return memo[total]
	}

	minCount := math.MaxInt32
	for i := 0; i < len(values); i++ {
		if total < values[i] {
			continue
		}
		rest := total - values[i]
		restCount := getMinCoinCount02(rest, values)
		if restCount == -1 {
			continue
		}

		currentCount := restCount + 1
		if currentCount < minCount {
			minCount = currentCount
		}
	}
	if minCount == math.MinInt32 {
		memo[total] = -1
		return -1
	}
	memo[total] = minCount
	return minCount
}

/* 动态规划 */
func getMinCoinCount04(total int, values []int) int {
	dp := make([]int, total + 1)
	for i := 0; i < total+1; i++ {
		dp[i] = total + 1
	}
	dp[0] = 0
	for i := 0; i < total+1; i++ {
		for j := 0; j < len(values); j++ {
			if i - values[j] < 0 {
				continue
			}
			dp[i] = min(dp[i - values[j]] + 1, dp[i])
		}
	}

	if dp[total] == total + 1 {
		return -1
	}
	return dp[total]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

