package _0050

// 2020.05.03
// https://leetcode-cn.com/problems/powx-n/comments/
func myPow(x float64, n int) float64 {
	var pow func(x float64, n int) float64
	if n < 0 {
		x = 1 / x
		n = -n
	}

	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		r := pow(x, n/2)
		if n%2 == 0 {
			return r * r
		} else {
			return r * r * x
		}
	}

	return pow(x, n)
}
