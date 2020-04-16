package _59

// 2020.04.14
// https://leetcode-cn.com/problems/spiral-matrix-ii/

func generateMatrix(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	var (
		matrix     = make([][]int, n)
		directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		dIdx       = 0
		idx        = 1
		total      = n * n
		cango      = func(r, c int) bool { return r >= 0 && r < n && c >= 0 && c < n && matrix[r][c] == 0 }
		nextD      func(r, c int) (nr, nc int)
		r, c       int
	)
	for idx, _ := range matrix {
		matrix[idx] = make([]int, n)
	}
	nextD = func(r, c int) (nr, nc int) {
		for {
			nr, nc = r+directions[dIdx][0], c+directions[dIdx][1]
			if !cango(nr, nc) {
				dIdx = (dIdx + 1) % 4
			} else {
				break
			}
		}

		return
	}

	for {
		matrix[r][c] = idx
		if idx == total {
			break
		}
		idx++
		r, c = nextD(r, c)
	}

	return matrix
}
