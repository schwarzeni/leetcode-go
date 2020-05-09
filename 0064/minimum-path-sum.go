package _0064

// 2020.05.09
// https://leetcode-cn.com/problems/minimum-path-sum/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	cache := make([][]int, len(grid))
	for idx, _ := range cache {
		cache[idx] = make([]int, len(grid[0]))
	}
	for r, cc := range grid {
		for c, v := range cc {
			if r == 0 && c == 0 {
				cache[r][c] = v
			} else {
				if c == 0 || (r != 0 && cache[r][c-1] >= cache[r-1][c]) {
					cache[r][c] = v + cache[r-1][c]
				}
				if r == 0 || (c != 0 && cache[r][c-1] < cache[r-1][c]) {
					cache[r][c] = v + cache[r][c-1]
				}
			}
		}
	}

	return cache[len(grid)-1][len(grid[0])-1]
}
