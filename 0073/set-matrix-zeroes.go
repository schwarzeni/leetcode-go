package _0073

// 2020.05.16
// https://leetcode-cn.com/problems/set-matrix-zeroes/
//func setZeroes(matrix [][]int) {
//	rc := len(matrix)
//	if rc == 0 {
//		return
//	}
//	cc := len(matrix[0])
//	visited := make([][]bool, rc)
//	for idx, _ := range visited {
//		visited[idx] = make([]bool, cc)
//	}
//	setZero := func(currR, currC int) {
//		for r := 0; r < rc; r++ {
//			if matrix[r][currC] != 0 {
//				visited[r][currC] = true
//				matrix[r][currC] = 0
//			}
//		}
//		for c := 0; c < cc; c++ {
//			if matrix[currR][c] != 0 {
//				visited[currR][c] = true
//				matrix[currR][c] = 0
//			}
//		}
//	}
//	for r := 0; r < rc; r++ {
//		for c := 0; c < cc; c++ {
//			if matrix[r][c] == 0 && !visited[r][c] {
//				setZero(r, c)
//			}
//		}
//	}
//}

func setZeroes(matrix [][]int) {
	rc := len(matrix)
	if rc == 0 {
		return
	}
	cc := len(matrix[0])
	r0_flag, c0_flag := false, false

	for r := 0; r < rc; r++ {
		if matrix[r][0] == 0 {
			r0_flag = true
			break
		}
	}
	for c := 0; c < cc; c++ {
		if matrix[0][c] == 0 {
			c0_flag = true
			break
		}
	}
	for r := 1; r < rc; r++ {
		for c := 1; c < cc; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c], matrix[r][0] = 0, 0
			}
		}
	}
	for r := 1; r < rc; r++ {
		if matrix[r][0] == 0 {
			for c := 1; c < cc; c++ {
				matrix[r][c] = 0
			}
		}
	}
	for c := 1; c < cc; c++ {
		if matrix[0][c] == 0 {
			for r := 1; r < rc; r++ {
				matrix[r][c] = 0
			}
		}
	}
	if c0_flag {
		for c := 0; c < cc; c++ {
			matrix[0][c] = 0
		}
	}
	if r0_flag {
		for r := 0; r < rc; r++ {
			matrix[r][0] = 0
		}
	}
}
