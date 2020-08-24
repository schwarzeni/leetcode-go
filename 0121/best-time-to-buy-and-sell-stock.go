package _0121

// 2020.08.24
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	ans := 0
	for _, v := range prices[1:] {
		if earn := v - min; earn > ans {
			ans = earn
		}
		if v < min {
			min = v
		}
	}
	return ans
}
