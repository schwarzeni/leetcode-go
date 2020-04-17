package _0007

import "math"

// 2020.04.17
// https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	bound := math.MaxInt32
	if x < 0 {
		bound = math.MinInt32
	}
	r := 0

	for x != 0 {
		i := x % 10
		x = x / 10
		if isPlusOverflow(r, 10, bound) || isAddOverflow(r*10, i, bound) {
			return 0
		}
		r = r*10 + i
	}

	return r
}

func isAddOverflow(raw, incre, bound int) bool {
	if bound > 0 {
		return bound-raw < incre
	}
	return bound-raw > incre
}

func isPlusOverflow(raw, plus, bound int) bool {
	if bound > 0 {
		return bound/plus < raw
	}
	return bound/plus > raw
}
