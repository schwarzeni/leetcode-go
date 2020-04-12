package _29

import "math"

// 2020.04.12
// https://leetcode-cn.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 { // 处理溢出情况
		return math.MaxInt32
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		} else {
			return 0
		}
	}

	var (
		dd     = dividend
		dr     = divisor
		result int
	)
	if dd > 0 {
		if divisor < 0 {
			dr = -dr
		}
		for dd-dr >= 0 {
			x := 0
			for dd-(dr<<1<<x) >= 0 {
				x += 1
			}
			result += 1 << x
			dd -= dr << x
		}
		if divisor < 0 {
			result = -result
		}
		return result
	}

	if dd < 0 {
		if divisor > 0 {
			dr = -dr
		}
		for dd-dr <= 0 {
			x := 0
			for dd-(dr<<1<<x) <= 0 {
				x += 1
			}
			result += 1 << x
			dd -= dr << x
		}
		if divisor > 0 {
			result = -result
		}
		return result
	}

	return 0
}

// 23 / 4
//

// a * b
// a * 2 ==> a << 2
// x * a = b
// 10 = 1 << 3 + 2
