package _0221

// 2020.05.08
// https://leetcode-cn.com/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	cache := make([][]int, len(matrix))
	for idx, _ := range cache {
		cache[idx] = make([]int, len(matrix[0]))
	}
	maxEdge := 0
	for r := 0; r < len(cache); r++ {
		for c := 0; c < len(cache[0]); c++ {
			if matrix[r][c] == '1' {
				cache[r][c] = 1
				if r != 0 && c != 0 {
					cache[r][c] = min(cache[r-1][c], cache[r][c-1], cache[r-1][c-1]) + 1
					if cache[r][c] > maxEdge {
						maxEdge = cache[r][c]
					}
				} else if maxEdge == 0 {
					maxEdge = 1
				}
			}
		}
	}
	return maxEdge * maxEdge
}

func min(args ...int) (r int) {
	r = args[0]
	for _, v := range args[1:] {
		if v < r {
			r = v
		}
	}
	return
}
