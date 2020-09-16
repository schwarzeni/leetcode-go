package _0413

// 2020.09.16
// https://leetcode-cn.com/problems/arithmetic-slices/
func numberOfArithmeticSlices(A []int) int {
	l := len(A)
	if l < 3 {
		return 0
	}
	cacheN := make([]bool, l)
	count := 0
	for i := 2; i < l; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			count++
			cacheN[i] = true
		}
	}
	for ll := 3; ll < l; ll++ { // ll: 等差数列的长度-1，同时也为数组起始idx
		cacheN_1 := cacheN
		cacheN = make([]bool, l)
		for i := ll; i < l; i++ {
			if A[i]-A[i-1] == A[i-1]-A[i-2] && cacheN_1[i-1] {
				count++
				cacheN[i] = true
			}
		}
	}
	return count
}
