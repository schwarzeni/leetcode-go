package _0983

// 2020.05.06
// https://leetcode-cn.com/problems/minimum-cost-for-tickets/

//func mincostTickets(days []int, costs []int) int {
//	var iter func(d int, idx int, cost int) int
//	var cache = [366][366]map[int]int{}
//	for i := 0; i < 366; i++ {
//		for j := 0; j < 366; j++ {
//			cache[i][j] = make(map[int]int)
//		}
//	}
//
//	iter = func(d int, idx int, cost int) int {
//		if d > 365 || d > days[len(days)-1] {
//			return cost
//		}
//		if c, ok := cache[d][idx][cost]; ok {
//			return c
//		}
//		for days[idx] < d {
//			idx++
//		}
//		needToVisit := false
//		if d == days[idx] {
//			needToVisit = true
//			idx++
//			if idx == len(days) {
//				return cost + costs[0]
//			}
//		}
//		c := math.MaxInt64
//		if needToVisit {
//			if cc := iter(d+1, idx, cost+costs[0]); cc != -1 {
//				c = cc
//			}
//		} else {
//			if cc := iter(d+1, idx, cost); cc != -1 {
//				c = cc
//			}
//		}
//		if cc := iter(d+7, idx, cost+costs[1]); cc != -1 && c > cc {
//			c = cc
//		}
//		if cc := iter(d+30, idx, cost+costs[2]); cc != -1 && c > cc {
//			c = cc
//		}
//		cache[d][idx][cost] = c
//		return c
//	}
//
//	return iter(1, 0, 0)
//}

func mincostTickets(days []int, costs []int) int {
	n := days[len(days)-1]
	dp := make([]int, n+1)
	for _, v := range days {
		dp[v] = -1
	}
	for i := 1; i <= n; i++ {
		if dp[i] == 0 {
			dp[i] = dp[i-1]
		} else {
			a := dp[i-1] + costs[0]
			b := costs[1]
			c := costs[2]
			if i-7 >= 0 {
				b += dp[i-7]
			}
			if i-30 >= 0 {
				c += dp[i-30]
			}
			dp[i] = min(a, b, c)
		}
	}
	return dp[n]
}

func min(arr ...int) int {
	min := arr[0]
	for _, v := range arr[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
