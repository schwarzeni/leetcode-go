package _0037

// 2020.04.28
// https://leetcode-cn.com/problems/sudoku-solver/
//func solveSudoku(board [][]byte) {
//	const (
//		T_VISITED   = 1
//		T_UNVISITED = 0
//		T_STATIC    = 2
//	)
//	var (
//		rMap    = make([][]int, 10)
//		cMap    = make([][]int, 10)
//		gMap    = make([][]int, 10)
//		dfs     func(r, c int) bool
//		changeT = func(r, c, v, t int) {
//			rMap[r][v] = t
//			cMap[c][v] = t
//			gMap[gpos(r, c)][v] = t
//		}
//		count int
//	)
//	for idx, _ := range rMap {
//		rMap[idx] = make([]int, 10)
//		cMap[idx] = make([]int, 10)
//		gMap[idx] = make([]int, 10)
//	}
//	for r, cs := range board {
//		for c, v := range cs {
//			if v != '.' {
//				changeT(r, c, int(v)-48, T_STATIC)
//			} else {
//				count++
//			}
//		}
//	}
//
//	dfs = func(rr, cc int) bool {
//		r, c := rr, cc
//
//		// 遍历正方形时有点想当然了，在这里卡了好久
//		for r < 9 && c < 9 {
//			if board[r][c] == '.' {
//				break
//			}
//			if c == 8 {
//				r++
//				c = 0
//			} else {
//				c++
//			}
//		}
//
//		if (r == 9 || c == 9) && count > 0 {
//			return false
//		}
//		for i := 1; i <= 9; i++ {
//			if rMap[r][i] == T_UNVISITED && cMap[c][i] == T_UNVISITED && gMap[gpos(r, c)][i] == T_UNVISITED {
//				changeT(r, c, i, T_VISITED)
//				board[r][c] = byte(i + 48)
//				count--
//				var result bool
//				if (c == 8 && r == 8) || count == 0 {
//					return true
//				} else if c == 8 && r != 8 {
//					result = dfs(r+1, 0)
//				} else if c != 8 {
//					result = dfs(r, c+1)
//				}
//				if !result {
//					board[r][c] = '.'
//					changeT(r, c, i, T_UNVISITED)
//					count++
//				} else {
//					return true
//				}
//			}
//		}
//		return false
//	}
//
//	dfs(0, 0)
//}

func solveSudoku(board [][]byte) {
	const (
		T_VISITED   = 1
		T_UNVISITED = 0
		T_STATIC    = 2
	)
	var (
		rMap    = make([][]int, 10)
		cMap    = make([][]int, 10)
		gMap    = make([][]int, 10)
		dfs     func(r, c int) bool
		changeT = func(r, c, v, t int) { rMap[r][v], cMap[c][v], gMap[(r/3)*3+c/3][v] = t, t, t }
	)
	for idx, _ := range rMap {
		rMap[idx] = make([]int, 10)
		cMap[idx] = make([]int, 10)
		gMap[idx] = make([]int, 10)
	}
	for r, cs := range board {
		for c, v := range cs {
			if v != '.' {
				changeT(r, c, int(v)-48, T_STATIC)
			}
		}
	}

	dfs = func(r, c int) bool {
		if r == 8 && c == 9 {
			return true
		} else if c == 9 {
			r++
			c = 0
		}

		if board[r][c] == '.' {
			for i := 1; i <= 9; i++ {
				if rMap[r][i] == T_UNVISITED && cMap[c][i] == T_UNVISITED && gMap[(r/3)*3+c/3][i] == T_UNVISITED {
					changeT(r, c, i, T_VISITED)
					board[r][c] = byte(i + 48)
					if (r == 8 && c == 8) || dfs(r, c+1) {
						return true
					}
					board[r][c] = '.'
					changeT(r, c, i, T_UNVISITED)
				}
			}
			return false
		} else {
			return dfs(r, c+1)
		}
	}

	dfs(0, 0)
}

// 5 3 4 6 7 8 9 1 2
//6 7 2 1 9 5 3 4 8
//1 9 8 3 4 2 5 6 7
//8 5 9 7 6 1 4 2 3
