package _60

// 2020.04.05
// https://leetcode-cn.com/problems/permutation-sequence/
func getPermutation(n int, k int) string {
	var (
		mark           = make([]bool, n)
		idx            = 1
		result         = make([]byte, n)
		subPermutation func(level int) bool
	)

	subPermutation = func(level int) bool {
		if idx > k {
			return true
		}
		for i := 0; i < n; i++ {
			if !mark[i] {
				mark[i] = true
				result[level] = byte(i + 49)
				if level == n-1 {
					idx += 1
				}
				if subPermutation(level + 1) {
					return true
				}
				mark[i] = false
			}
		}
		return false
	}
	if n == 1 {
		return "1"
	}
	_ = subPermutation(0)
	return string(result)
}

// 1 2 3 4
