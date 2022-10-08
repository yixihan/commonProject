package main

import "math"

// 动态规划 前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return dp[len(prices)-1]
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
