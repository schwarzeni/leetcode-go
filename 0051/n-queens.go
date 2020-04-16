package _51

// 2020.04.04
// https://leetcode-cn.com/problems/n-queens/
func solveNQueens(n int) [][]string {
	const (
		T_QUEEN = 'Q'
		T_BLANK = '.'
	)
	var (
		result [][]string
		grid   = make([][]byte, n)
		marked = make([][]bool, 4) // 行、列、k=1, k=-1
		cango  = func(r, c int) bool {
			return !marked[0][r] && !marked[1][c] && !marked[2][r-c+n] && !marked[3][c+r]
		}
		setGridStatus = func(r, c int, status bool) {
			marked[0][r] = status
			marked[1][c] = status
			marked[2][r-c+n] = status
			marked[3][c+r] = status
		}
		getResult = func() {
			r := make([]string, n)
			for idx, _ := range r {
				r[idx] = string(grid[idx])
			}
			result = append(result, r)
		}

		checkLevel func(l int)
	)
	for idx, _ := range grid {
		grid[idx] = make([]byte, n)
		for i, _ := range grid[idx] {
			grid[idx][i] = T_BLANK
		}
	}
	for idx, _ := range marked {
		marked[idx] = make([]bool, n*2)
	}

	checkLevel = func(l int) {
		for c := 0; c < n; c++ {
			if cango(l, c) {
				setGridStatus(l, c, true)
				grid[l][c] = T_QUEEN
				if l == n-1 {
					getResult()
				} else {
					checkLevel(l + 1)
				}
				grid[l][c] = T_BLANK
				setGridStatus(l, c, false)
			}
		}
	}

	if n == 0 {
		return [][]string{{}}
	}
	checkLevel(0)
	return result
}
