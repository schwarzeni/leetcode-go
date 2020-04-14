package _48

// 2020.04.14
// https://leetcode-cn.com/problems/rotate-image/
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	var (
		r    = 0
		c    = 0
		gLen = len(matrix)
	)

	if len(matrix) != len(matrix[0]) {
		return
	}

	for gLen > 1 {
		// do something
		minR := r
		maxR := r + gLen - 1
		minC := c
		maxC := maxR
		rr := 0
		for {
			matrix[minR][minC+rr], matrix[minR+rr][maxC], matrix[maxR][maxC-rr], matrix[maxR-rr][minC] =
				matrix[maxR-rr][minC], matrix[minR][minC+rr], matrix[minR+rr][maxC], matrix[maxR][maxC-rr]
			if rr += 1; minR+rr >= maxR {
				break
			}
		}

		r += 1
		c += 1
		gLen = gLen - 2
	}
}

// 0  1  2  3
// 4  5  6  7
// 8  9  10 11
// 12 13 14 15

// 12 8  4  0
// 13 9  5  1
// 14 10 6  2
// 15 11 7  3
