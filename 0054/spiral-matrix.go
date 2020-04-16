package _54

// 2020.04.14
// https://leetcode-cn.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var (
		r      = 0
		c      = 0
		rLen   = len(matrix)
		cLen   = len(matrix[0])
		result = make([]int, rLen*cLen)
		idx    = 0
	)

	for rLen > 0 && cLen > 0 {
		minR := r
		minC := c
		maxR := r + rLen - 1
		maxC := c + cLen - 1
		i := 0

		for i = minC; i <= maxC; i++ {
			result[idx] = matrix[minR][i]
			idx++
		}

		for i = minR + 1; i <= maxR; i++ {
			result[idx] = matrix[i][maxC]
			idx++
		}

		if maxR == minR || maxC == minC {
			break
		}

		for i = maxC - 1; i >= minC; i-- {
			result[idx] = matrix[maxR][i]
			idx++
		}

		for i = maxR - 1; i > minR; i-- {
			result[idx] = matrix[i][minC]
			idx++
		}

		r++
		c++
		rLen -= 2
		cLen -= 2
	}

	return result
}
