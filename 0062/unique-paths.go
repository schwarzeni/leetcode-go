package _0062

// 2020.05.07
// https://leetcode-cn.com/problems/unique-paths/
//func uniquePaths(m int, n int) int {
//	var (
//		visit       = make([][]int, n)
//		directions  = [][]int{{0, 1}, {1, 0}}
//		cango       = func(r, c int) bool { return r >= 0 && r < n && c >= 0 && c < m }
//		dfs         func(r, c int)
//		count       int
//		S_UNVISITED = 0
//		S_VISITED   = 1
//	)
//	for idx, _ := range visit {
//		visit[idx] = make([]int, m)
//	}
//
//	dfs = func(r, c int) {
//		if r == n-1 && c == m-1 {
//			count++
//			return
//		}
//		for _, d := range directions {
//			nr, nc := r+d[0], c+d[1]
//			if cango(nr, nc) && visit[nr][nc] == S_UNVISITED {
//				visit[nr][nc] = S_VISITED
//				dfs(nr, nc)
//				visit[nr][nc] = S_UNVISITED
//			}
//		}
//	}
//
//	visit[0][0] = S_VISITED
//	dfs(0, 0)
//	return count
//}

func uniquePaths(m int, n int) int {
	var count = make([][]int, n)

	for idx, _ := range count {
		count[idx] = make([]int, m)
	}
	for idx := 0; idx < m; idx++ {
		count[0][idx] = 1
	}
	for idx := 0; idx < n; idx++ {
		count[idx][0] = 1
	}
	for r := 1; r < n; r++ {
		for c := 1; c < m; c++ {
			count[r][c] = count[r-1][c] + count[r][c-1]
		}
	}

	return count[n-1][m-1]
}
