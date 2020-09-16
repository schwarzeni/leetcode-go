package _0446

// 2020.09.16
// https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/
func numberOfArithmeticSlices(A []int) int {
	l := len(A)
	if l < 3 {
		return 0
	}
	cache := make([]map[int]int, l)
	for i := range cache {
		cache[i] = make(map[int]int, l)
	}
	count := 0
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]
			cache[i][diff] += cache[j][diff] + 1
			if v, ok := cache[j][diff]; ok {
				count += v
			}
		}
	}
	return count
}
