package _0096

// 2020.06.11
// https://leetcode-cn.com/problems/unique-binary-search-trees/

func numTrees(n int) int {
	var numTrees1 func(n int) int
	cache := make(map[int]int)
	cache[1], cache[0] = 1, 1
	numTrees1 = func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		var res int
		for i := 1; i <= n; i++ {
			res += numTrees1(i-1) * numTrees1(n-i)
		}
		cache[n] = res
		return res
	}
	return numTrees1(n)
}
