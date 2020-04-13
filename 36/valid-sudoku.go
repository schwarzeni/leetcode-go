package _36

// 2020.04.13
// https://leetcode-cn.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	const T_BLANK = '.'
	marked := make([][9]map[byte]struct{}, 3, 3)
	for idx, _ := range marked {
		for i, _ := range marked[idx] {
			marked[idx][i] = make(map[byte]struct{})
		}
	}

	for r, cs := range board {
		for c, v := range cs {
			if v != T_BLANK {
				for idx, pos := range getPos(r, c) {
					if _, ok := marked[idx][pos][v]; ok {
						return false
					}
					marked[idx][pos][v] = struct{}{}
				}
			}
		}
	}

	return true
}

func getPos(r, c int) (pos []int) {
	pos = make([]int, 3, 3)
	pos[0] = r
	pos[1] = c
	pos[2] = (r/3)*3 + c/3
	return
}

// row  0 1 2 ... 8
// col  0 1 2 ... 8
// grid 0 1 2 ... 8

// (4, 4)
// 1, 1 => 1*3+1
