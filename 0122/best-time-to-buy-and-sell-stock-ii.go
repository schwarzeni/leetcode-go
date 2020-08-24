package _0122

// 2020.08.24
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp[i][0]表示第i天不持有股票
	// dp[i][1]表示第i天持有股票
	dp := make([][2]int, len(prices))

	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
